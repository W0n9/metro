package data

// MetroData 代表地铁数据的完整结构
type MetroData struct {
	StaDict    map[string]Station  `json:"staDict"`
	StaToId    map[string]int      `json:"staToId"`
	FreeDis    int                 `json:"freeDis"`
	LineDetail map[string]LineInfo `json:"lineDetail"`
}

// Station 代表单个站点的信息
type Station struct {
	Name  string            `json:"name"`
	Edges map[string][]Edge `json:"edges"`
}

// Edge 代表站点之间的连接边
type Edge struct {
	Dis int `json:"dis"`
}

// LineInfo 代表地铁线路信息
type LineInfo struct {
	Name    string `json:"name"`
	StaList []int  `json:"staList"`
}

// SearchResult 代表搜索结果
type SearchResult struct {
	Distance int
	Price    int
	Path     []string
}
