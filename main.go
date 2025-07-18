package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"metro-go/src/utils"
)

func main() {
	// 检查命令行参数
	if len(os.Args) != 3 {
		fmt.Println("用法: go run main.go <出发站名> <车费预算(元)>")
		return
	}

	startName := strings.TrimSpace(os.Args[1])
	budgetStr := os.Args[2]

	budget, err := strconv.Atoi(budgetStr)
	if err != nil {
		fmt.Println("预算必须为整数！")
		return
	}

	fmt.Println("正在加载地铁数据...")

	// 加载地铁数据
	data, err := utils.LoadMetroData()
	if err != nil {
		fmt.Printf("加载数据失败: %v\n", err)
		return
	}

	// 获取起始站点ID
	startID := utils.NameToID(data.StaToId, startName)
	if startID == "-1" {
		fmt.Println("站名不存在！")
		return
	}

	fmt.Printf("以 %s 为起点，预算 %d 元，可达站点如下：\n", startName, budget)

	// 执行Dijkstra算法搜索
	result := utils.DijkstraAll(data.StaDict, startID, data.FreeDis)

	count := 0
	for stationID, info := range result {
		if info.Price == budget && stationID != startID {
			// 转换路径ID为站名
			pathNames := make([]string, len(info.Path))
			for i, pathID := range info.Path {
				pathNames[i] = utils.IDToName(data.StaDict, pathID)
			}

			// 获取站点所属线路
			lines := utils.GetStationLines(stationID, data.LineDetail)
			lineStr := "未知"
			if len(lines) > 0 {
				lineStr = strings.Join(lines, ", ")
			}

			stationName := utils.IDToName(data.StaDict, stationID)
			pathStr := strings.Join(pathNames, " -> ")

			fmt.Printf("%s | 距离: %dm | 票价: %d元 | 路径: %s | 线路: %s\n",
				stationName, info.Distance, info.Price, pathStr, lineStr)
			count++
		}
	}

	fmt.Printf("共 %d 个站点可达。\n", count)
}
