package tree

import (
	"testing"
)

func TestCalcShannoEnt(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
	}
	CalcShannoEnt(&dataSet)
}
