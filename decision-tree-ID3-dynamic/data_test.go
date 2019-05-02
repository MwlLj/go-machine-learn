package tree

import (
	// "log"
	"fmt"
	"testing"
)

var _ = fmt.Println

func TestSplitDataSet(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
	}
	value := "0"
	vec := SplitDataSet(&dataSet, 0, &value)
	printDataSet(vec)
	value = "1"
	vec = SplitDataSet(&dataSet, 0, &value)
	printDataSet(vec)
}

func TestFindGreaterThanFeatureIndexSet(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
		{"0", "1", "2", "no"},
	}
	fmt.Println("FindGreaterThanFeatureIndexSet start")
	vec := FindGreaterThanFeatureIndexSet(&dataSet, 1)
	printDataSet(vec)
	vec = FindGreaterThanFeatureIndexSet(&dataSet, 2)
	printDataSet(vec)
	fmt.Println("FindGreaterThanFeatureIndexSet end")
}

func TestSplitDataSetByFeatureNumber(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
		{"0", "1", "2", "no"},
	}
	fmt.Println("SplitDataSetByFeatureNumber start")
	subDataSet := SplitDataSetByFeatureNumber(&dataSet)
	for _, item := range *subDataSet {
		fmt.Println(item.FeatureNumberMax, *item.DataSet)
	}
}
