package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"-"` // 忽略该字段
	Subjects []string `json:"sub"`
	IsOk     bool     `json:"ok,string"`
	Price    float64  `json:"price,string,omitempty"`
}

func main() {
	it := IT{"", []string{"1", "2"}, true, 1.2}
	bytes, err := json.Marshal(it)
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	jsonStr := string(bytes)
	fmt.Println(jsonStr)

	var it2 IT
	json.Unmarshal([]byte(jsonStr), &it2)
	fmt.Println(it2)

}
