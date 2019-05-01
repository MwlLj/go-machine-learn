package tree

import (
	"bytes"
	"encoding/json"
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

	lables := []string{"no surfacing", "flippers"}
	node := CreateTree(&dataSet, &lables)
	b, err := json.Marshal(node)
	if err != nil {
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return
	}
	fmt.Println(out.String())

	featureValues := []string{"1", "1"}
	classify := FindByOrderFeature(node, &featureValues)
	if classify != nil {
		fmt.Println(*classify)
	}
}
