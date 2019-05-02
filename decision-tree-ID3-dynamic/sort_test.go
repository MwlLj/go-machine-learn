package tree

import (
	"fmt"
	"testing"
)

func TestClassifyTimesSort(t *testing.T) {
	values := []string{
		"dog", "dog", "dog", "cat", "cat",
	}
	sorts := classifyTimesSort(&values)
	fmt.Println(*sorts)
}

func TestFeatureNumberMax(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
		{"0", "1", "2", "no"},
	}
	max, _ := featureNumberMax(&dataSet)
	fmt.Println("max", max)
}
