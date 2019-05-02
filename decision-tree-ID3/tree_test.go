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
		{"1", "0", "yes"},
		{"1", "0", "no"},
		{"0", "2", "no"},
		{"0", "2", "no"},
		{"0", "3", "no"},
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
	classify := FindByOrderFeature(node, &featureValues, &lables)
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

	featureValues := []string{"pre", "myope", "yes", "normal"}
	classify := FindByOrderFeature(node, &featureValues, &lables)
	if classify != nil {
		fmt.Println(*classify)
	}
}
