package tree

import (
	"errors"
	"fmt"
)

var _ = fmt.Println
var _ = errors.New

type Node struct {
	Classify     string  `json:"classify"`
	FeatureList  *[]Node `json:"feature-list"`
	Lable        string  `json:"lable"`
	FeatureValue string  `json:"feature-value"`
}

func ChooseBestFeature(dataSet *[][]string, featureNumber int) int {
	baseEnt := CalcShannoEnt(dataSet)
	setLen := len(*dataSet)
	bestInfoGain := 0.0
	bestFeature := 0
	for i := 0; i < featureNumber; i++ {
		values, _ := getFeatureValueSet(dataSet, i)
		featureEnt := 0.0
		for _, item := range *values {
			itemTmp := item
			subSet := SplitDataSet(dataSet, i, &itemTmp)
			subLen := len(*subSet)
			featureEnt += float64(subLen) / float64(setLen) * CalcShannoEnt(subSet)
		}
		infoGain := baseEnt - featureEnt
		if infoGain > bestInfoGain {
			bestInfoGain = infoGain
			bestFeature = i
		}
	}
	return bestFeature
}

func CreateTree(dataSet *[][]string, lables *[]string) *Node {
	featureNumber, _ := featureNumberMax(dataSet)
	classifySet, classifys := getFeatureValueSet(rightfulSet, featureNumber)
	if featureNumber == 0 {
		// no feature
		pairList := classifyTimesSort(classifys)
		node := Node{
			Classify:    (*pairList)[0].Key.(string),
			FeatureList: nil,
			Lable:       "",
		}
		return &node
	}
	if len(*classifySet) == 1 {
		node := Node{
			Classify:    (*classifySet)[0],
			FeatureList: nil,
			Lable:       "",
		}
		return &node
	}
	bestFeatureIndex := ChooseBestFeature(rightfulSet)
	// name := (*lables)[bestFeatureIndex]
	sets, _ := getFeatureValueSet(rightfulSet, bestFeatureIndex)
	lable := (*lables)[bestFeatureIndex]
	nodes := []Node{}
	for _, item := range *sets {
		itemTmp := item
		subSet := SplitDataSet(rightfulSet, bestFeatureIndex, &itemTmp)
		node := CreateTree(subSet, removeArrayIndex(lables, bestFeatureIndex))
		node.FeatureValue = item
		nodes = append(nodes, *node)
	}
	return &Node{
		Classify:    "",
		FeatureList: &nodes,
		Lable:       lable,
	}
}

func FindByOrderFeature(node *Node, featureValues *[]string) *string {
	if node.FeatureList == nil {
		return &node.Classify
	}
	for _, item := range *node.FeatureList {
		first := (*featureValues)[0]
		after := (*featureValues)[1:]
		if first == item.FeatureValue && item.FeatureList == nil {
			return &item.Classify
		}
		if first == item.FeatureValue && item.FeatureList != nil {
			return FindByOrderFeature(&item, &after)
		}
	}
	return nil
}

func getFeatureValueSet(dataSet *[][]string, featureIndex int) (featureSet *[]string, features *[]string) {
	sets := []string{}
	vecs := []string{}
	mapTmp := map[string]bool{}
	for _, item := range *dataSet {
		value := item[featureIndex]
		if _, ok := mapTmp[value]; !ok {
			sets = append(sets, value)
			mapTmp[value] = true
		}
		vecs = append(vecs, value)
	}
	return &sets, &vecs
}

func removeArrayIndex(arr *[]string, index int) *[]string {
	retArr := (*arr)[0:index]
	retArr = append(retArr, (*arr)[index+1:]...)
	return &retArr
}
