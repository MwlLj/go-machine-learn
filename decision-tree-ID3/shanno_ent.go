package tree

import (
	"fmt"
	"math"
)

var _ = fmt.Println

func CalcShannoEnt(dataSet *[][]string) float64 {
	mapTmp := map[string]int{}
	for _, item := range *dataSet {
		v := item[len(item)-1]
		_, ok := mapTmp[v]
		if !ok {
			mapTmp[v] = 0
		}
		mapTmp[v] += 1
	}
	shannoEnt := 0.0
	number := len(*dataSet)
	for _, value := range mapTmp {
		prob := float64(value) / float64(number)
		shannoEnt -= prob * math.Log2(prob)
	}
	return shannoEnt
}
