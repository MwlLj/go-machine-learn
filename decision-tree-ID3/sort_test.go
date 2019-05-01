package tree

import (
	"fmt"
	"testing"
)

func TestTimesSort(t *testing.T) {
	values := []string{
		"dog", "dog", "dog", "cat", "cat",
	}
	sorts := timesSort(&values)
	fmt.Println(*sorts)
}
