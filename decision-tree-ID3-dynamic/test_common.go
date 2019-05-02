package tree

import (
	"bytes"
	"fmt"
)

func printDataSet(dataSet *[][]string) {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	j := 0
	leng := len(*dataSet)
	for _, item := range *dataSet {
		j += 1
		buf.WriteString("[")
		i := 0
		length := len(item)
		for _, it := range item {
			i += 1
			buf.WriteString(it)
			if i < length {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("]")
		if j < leng {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	fmt.Println(buf.String())
}
