package tree

import (
	"fmt"
	"testing"
)

func TestChooseBestFeature(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
	}
	bestFeature := ChooseBestFeature(&dataSet)
	fmt.Println(bestFeature)
}
