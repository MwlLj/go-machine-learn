package tree

import (
	// "fmt"
	"testing"
)

type Float float64

func (this *Float) LessEqual(value interface{}) bool {
	if *this <= value.(Float) {
		return true
	}
	return false
}

func TestBinaryDataSetSplit(t *testing.T) {
	dataSet := [][]Float{
		{Float(1.0), Float(0.0), Float(0.0), Float(0.0)},
		{Float(0.0), Float(1.0), Float(0.0), Float(0.0)},
		{Float(0.0), Float(0.0), Float(1.0), Float(0.0)},
		{Float(0.0), Float(0.0), Float(0.0), Float(1.0)},
	}
	var _ = dataSet
	f := Float(0.0)
	var value []IValue = []*Float{
		&f,
	}
	var _ = value
}
