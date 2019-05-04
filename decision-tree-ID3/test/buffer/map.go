package buffer

import (
	"sort"
)

type Map struct {
	Key   string
	Value string
}

func Map2Array(data *map[string]string) *[]Map {
	retVec := []Map{}
	keys := []string{}
	for _, key := range *data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		m := Map{
			Key:   key,
			Value: (*data)[key],
		}
		retVec = append(retVec, m)
	}
	return &retVec
}

func Array2Map(data *[]Map) *map[string]string {
	retMap := map[string]string{}
	for _, item := range *data {
		retMap[item.Key] = item.Value
	}
	return &retMap
}
