package utils

import (
	"container/heap"
	"metro-go/src/data"
)

// PriorityQueue 实现优先队列用于 Dijkstra 算法
type PriorityQueue []*Item

type Item struct {
	Distance  int
	StationID string
	Path      []string
	Index     int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// DijkstraAll 使用Dijkstra算法，返回所有可达站点及路径、距离、票价
func DijkstraAll(staDict map[string]data.Station, startID string, freeDis int) map[string]data.SearchResult {
	visited := make(map[string]bool)
	result := make(map[string]data.SearchResult)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 初始化起点
	heap.Push(&pq, &Item{
		Distance:  0,
		StationID: startID,
		Path:      []string{startID},
	})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)

		if visited[current.StationID] {
			continue
		}

		visited[current.StationID] = true
		price := CalcPrice(current.Distance, freeDis)

		// 复制路径
		path := make([]string, len(current.Path))
		copy(path, current.Path)

		result[current.StationID] = data.SearchResult{
			Distance: current.Distance,
			Price:    price,
			Path:     path,
		}

		// 访问邻居节点
		if station, exists := staDict[current.StationID]; exists {
			for neighborID, edges := range station.Edges {
				if !visited[neighborID] {
					for _, edge := range edges {
						nextDistance := current.Distance + edge.Dis
						nextPath := make([]string, len(current.Path))
						copy(nextPath, current.Path)
						nextPath = append(nextPath, neighborID)

						heap.Push(&pq, &Item{
							Distance:  nextDistance,
							StationID: neighborID,
							Path:      nextPath,
						})
					}
				}
			}
		}
	}

	return result
}
