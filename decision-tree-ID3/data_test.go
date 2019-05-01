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
