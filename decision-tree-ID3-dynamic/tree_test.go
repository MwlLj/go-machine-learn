package tree

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	// "strconv"
	"strings"
	"testing"
)

func TestChooseBestFeature(t *testing.T) {
	dataSet := [][]string{
		{"1", "1", "yes"},
		{"1", "1", "yes"},
		{"1", "0", "no"},
		{"0", "1", "no"},
		{"0", "1", "no"},
		{"young", "myope", "no", "reduced", "no lenses"},
		{"young", "myope", "no", "normal", "soft"},
		{"young", "myope", "yes", "reduced", "no lenses"},
		{"young", "myope", "yes", "normal", "hard"},
		{"young", "hyper", "no", "reduced", "no lenses"},
		{"young", "hyper", "no", "normal", "soft"},
		{"young", "hyper", "yes", "reduced", "no lenses"},
		{"young", "hyper", "yes", "normal", "hard"},
		{"pre", "myope", "no", "reduced", "no lenses"},
		{"pre", "myope", "no", "normal", "soft"},
		{"pre", "myope", "yes", "reduced", "no lenses"},
		{"pre", "myope", "yes", "normal", "hard"},
		{"pre", "hyper", "no", "reduced", "no lenses"},
		{"pre", "hyper", "no", "normal", "soft"},
		{"pre", "hyper", "yes", "reduced", "no lenses"},
		{"pre", "hyper", "yes", "normal", "no lenses"},
		{"presbyopic", "myope", "no", "reduced", "no lenses"},
		{"presbyopic", "myope", "no", "normal", "no lenses"},
		{"presbyopic", "myope", "yes", "reduced", "no lenses"},
		{"presbyopic", "myope", "yes", "normal", "hard"},
		{"presbyopic", "hyper", "no", "reduced", "no lenses"},
		{"presbyopic", "hyper", "no", "normal", "soft"},
		{"presbyopic", "hyper", "yes", "reduced", "no lenses"},
		{"presbyopic", "hyper", "yes", "normal", "no lenses"},
		{"presbyopic ", "hyper  ", "yes", "normal ", "no lenses"},
	}
	subDataSet := SplitDataSetByFeatureNumber(&dataSet)
	for _, sub := range *subDataSet {
		bestFeature := ChooseBestFeature(sub.DataSet, sub.FeatureNumberMax)
		fmt.Println("TestChooseBestFeature bestFeature: ", bestFeature)
	}

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

func TestChooseBestFeatureFromFile(t *testing.T) {
	file, err := os.Open("./dataset/lenses.txt")
	if err != nil {
		return
	}
	defer file.Close()
	dataSet := [][]string{}
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		fields := strings.Split(string(line), "\t")
		dataSet = append(dataSet, fields)
	}

	lables := []string{"age", "prescript", "astigmatic", "tearRate"}

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
}
