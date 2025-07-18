package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"metro-go/src/data"
)

const DataURL = "https://static.qinxr.cn/Hyacinth/farecalc.json"

// CalcPrice 计算票价，根据距离和免费距离计算
func CalcPrice(dis, freeDis int) int {
	dis -= freeDis
	if dis <= -freeDis {
		return 3
	} else if dis <= 6000 {
		return 3
	} else if dis <= 12000 {
		return 4
	} else if dis <= 22000 {
		return 5
	} else if dis <= 32000 {
		return 6
	} else {
		return int((dis-32000)/20000) + 7
	}
}

// LoadMetroData 从远程加载地铁数据
func LoadMetroData() (*data.MetroData, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(DataURL)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP错误: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	var metroData data.MetroData
	if err := json.Unmarshal(body, &metroData); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	return &metroData, nil
}

// NameToID 站名转ID
func NameToID(staToId map[string]int, name string) string {
	if id, exists := staToId[name]; exists {
		return strconv.Itoa(id)
	}
	return "-1"
}

// IDToName ID转站名
func IDToName(staDict map[string]data.Station, id string) string {
	if station, exists := staDict[id]; exists {
		return station.Name
	}
	return id
}

// GetStationLines 获取站点所属线路名
func GetStationLines(staID string, lineDetail map[string]data.LineInfo) []string {
	id, err := strconv.Atoi(staID)
	if err != nil {
		return []string{}
	}

	var lines []string
	for _, detail := range lineDetail {
		for _, stationID := range detail.StaList {
			if stationID == id {
				lines = append(lines, detail.Name)
				break
			}
		}
	}
	return lines
}
