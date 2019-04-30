package tree

import (
	"errors"
	"math"
)

type CShannoEnt struct {
}

func (this *CShannoEnt) CalcShannoEnt(dataSet []string) error {
	if len(dataSet) <= 1 {
		return errors.New("data is empty")
	}
}
