package tree

import (
	"errors"
)

func ChooseBestFeature(dataSet *[][]string) int {
	baseEnt := CalcShannoEnt(dataSet)
	bestEnt := 0.0
	bestFeature := -1
	setLen := len(*dataSet)
	featureNumber, _ := calcFeatureNumber(dataSet)
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
		if infoGain > bestEnt {
			bestEnt = featureEnt
			bestFeature = i
		}
	}
	return bestFeature
}

func CreateTree(dataSet *[][]string, lables *[]string) string {
	featureNumber, _ := calcFeatureNumber(dataSet)
	classifySet, classifys := getFeatureValueSet(dataSet, featureNumber+1)
	if featureNumber == 0 {
		// no feature
		pairList := timesSort(classifys)
		return (*pairList)[0].Key
	}
	if len(*classifySet) == 1 {
		return (*classifySet)[0]
	}
	return ""
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

func calcFeatureNumber(dataSet *[][]string) (int, error) {
	if len(*dataSet) < 1 {
		return 0, errors.New("data is empty")
	}
	featureNumber := len((*dataSet)[0]) - 1
	return featureNumber, nil
}
