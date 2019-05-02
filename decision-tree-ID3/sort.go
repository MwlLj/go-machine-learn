package tree

import (
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func timesSort(values *[]string) *PairList {
	mapTmp := map[string]int{}
	for _, item := range *values {
		if _, ok := mapTmp[item]; !ok {
			mapTmp[item] = 0
		}
		mapTmp[item] += 1
	}
	// sort by map value
	p := make(PairList, len(mapTmp))
	i := 0
	for k, v := range mapTmp {
		p[i] = Pair{k, v}
		i += 1
	}
	sort.Sort(p)
	return &p
}
