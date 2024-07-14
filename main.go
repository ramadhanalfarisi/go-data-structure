package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var arrayMap []map[string]any
	loopingCount := 1000
	for i := 1; i < loopingCount; i++ {
		arrayMap = append(arrayMap, map[string]any{"id": i, "name": fmt.Sprintf("name-%d", i)})
	}
	bst := CreateBST(arrayMap, "id")
	res := bst.Search(675)
	json, _ := json.Marshal(res)
	fmt.Println(string(json))
}