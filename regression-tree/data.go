package tree

import (
	"errors"
)

func BinaryDataSetSplit(dataSet *[][]IValue, featureIndex int, value IValue) (*[][]IValue, *[][]IValue, error) {
	dataSetLen := len(*dataSet)
	if dataSetLen == 0 {
		return nil, nil, errors.New("dataset is empty")
	}
	colLen := len((*dataSet)[0])
	if featureIndex > colLen-1 {
		return nil, nil, errors.New("index too larger")
	}
	leftSubSet := [][]IValue{}
	rightSubSet := [][]IValue{}
	for _, item := range *dataSet {
		feature := item[featureIndex]
		itemCopy := make([]IValue, colLen)
		copy(itemCopy, item)
		if feature.LessEqual(value) {
			// less equal
			rightSubSet = append(rightSubSet, itemCopy)
		} else {
			leftSubSet = append(leftSubSet, itemCopy)
		}
	}
	return &leftSubSet, &rightSubSet, nil
}
