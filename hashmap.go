package main

import "fmt"

// ArrayToMap converts a slice of maps to a map of maps, using the specified key to index the result.
// If the specified key does not exist in the input slice, an error is returned.
func ArraytoMap(array []map[string]any, key string) (map[string]map[string]any, error) {
	_, exist := array[0][key]
	if !exist {
		return nil, fmt.Errorf("ArrayToMap: key %s is not exist", key)
	}
	var result map[string]map[string]any = map[string]map[string]any{}
	for _, item := range array {
		keyString := fmt.Sprintf("%v", item[key])
		result[keyString] = item
	}
	return result, nil
}


