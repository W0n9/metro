package utils

import (
	"metro-go/src/data"
	"testing"
)

// TestCalcPrice 测试票价计算功能
func TestCalcPrice(t *testing.T) {
	freeDis := 0 // 设置免费距离为0，简化测试

	tests := []struct {
		distance int
		expected int
		desc     string
	}{
		{0, 3, "起点距离为0，票价应为3元"},
		{3000, 3, "3公里内，票价应为3元"},
		{6000, 3, "6公里，票价应为3元"},
		{6001, 4, "超过6公里，票价应为4元"},
		{12000, 4, "12公里，票价应为4元"},
		{12001, 5, "超过12公里，票价应为5元"},
		{22000, 5, "22公里，票价应为5元"},
		{22001, 6, "超过22公里，票价应为6元"},
		{32000, 6, "32公里，票价应为6元"},
		{32001, 7, "超过32公里，票价应为7元"},
		{52000, 8, "52公里，票价应为8元"},
		{52001, 8, "超过52公里，票价应为8元"},
	}

	for _, test := range tests {
		result := CalcPrice(test.distance, freeDis)
		if result != test.expected {
			t.Errorf("%s - 距离: %dm, 期望: %d元, 实际: %d元",
				test.desc, test.distance, test.expected, result)
		}
	}
}

// TestNameToID 测试站名转ID功能
func TestNameToID(t *testing.T) {
	staToId := map[string]int{
		"西直门": 123,
		"积水潭": 456,
		"东直门": 789,
	}

	tests := []struct {
		name     string
		expected string
		desc     string
	}{
		{"西直门", "123", "存在的站名应返回正确ID"},
		{"积水潭", "456", "存在的站名应返回正确ID"},
		{"不存在的站", "-1", "不存在的站名应返回-1"},
		{"", "-1", "空字符串应返回-1"},
	}

	for _, test := range tests {
		result := NameToID(staToId, test.name)
		if result != test.expected {
			t.Errorf("%s - 站名: '%s', 期望: %s, 实际: %s",
				test.desc, test.name, test.expected, result)
		}
	}
}

// TestIDToName 测试ID转站名功能
func TestIDToName(t *testing.T) {
	staDict := map[string]data.Station{
		"123": {Name: "西直门"},
		"456": {Name: "积水潭"},
		"789": {Name: "东直门"},
	}

	tests := []struct {
		id       string
		expected string
		desc     string
	}{
		{"123", "西直门", "存在的ID应返回正确站名"},
		{"456", "积水潭", "存在的ID应返回正确站名"},
		{"999", "999", "不存在的ID应返回原ID"},
		{"", "", "空字符串应返回空字符串"},
	}

	for _, test := range tests {
		result := IDToName(staDict, test.id)
		if result != test.expected {
			t.Errorf("%s - ID: '%s', 期望: %s, 实际: %s",
				test.desc, test.id, test.expected, result)
		}
	}
}

// TestGetStationLines 测试获取站点线路功能
func TestGetStationLines(t *testing.T) {
	lineDetail := map[string]data.LineInfo{
		"line1": {
			Name:    "一号线",
			StaList: []int{123, 456, 789},
		},
		"line2": {
			Name:    "二号线",
			StaList: []int{123, 999, 888},
		},
		"line3": {
			Name:    "三号线",
			StaList: []int{456, 777},
		},
	}

	tests := []struct {
		staID    string
		expected []string
		desc     string
	}{
		{"123", []string{"一号线", "二号线"}, "站点123应属于一号线和二号线"},
		{"456", []string{"一号线", "三号线"}, "站点456应属于一号线和三号线"},
		{"789", []string{"一号线"}, "站点789应只属于一号线"},
		{"999", []string{"二号线"}, "站点999应只属于二号线"},
		{"111", []string{}, "不存在的站点应返回空数组"},
		{"abc", []string{}, "无效ID应返回空数组"},
	}

	for _, test := range tests {
		result := GetStationLines(test.staID, lineDetail)

		// 比较数组长度
		if len(result) != len(test.expected) {
			t.Errorf("%s - 站点ID: %s, 期望线路数: %d, 实际线路数: %d",
				test.desc, test.staID, len(test.expected), len(result))
			continue
		}

		// 比较每个线路名（不考虑顺序）
		expectedMap := make(map[string]bool)
		for _, line := range test.expected {
			expectedMap[line] = true
		}

		for _, line := range result {
			if !expectedMap[line] {
				t.Errorf("%s - 站点ID: %s, 发现意外线路: %s",
					test.desc, test.staID, line)
			}
		}
	}
}
