package utils

import (
	"fmt"
	"testing"
)

func TestFindLable(t *testing.T) {
	dbLables := map[string][]string{}
	dbLables["1"] = []string{
		"young",
	}
	dbLables["2"] = []string{
		"young", "myope",
	}
	dbLables["3"] = []string{
		"young", "myope", "no",
	}
	dbLables["4"] = []string{
		"young", "myope", "no", "reduced",
	}
	{
		inputLable := []string{
			"young",
		}
		id, err := FindLable(&dbLables, &inputLable)
		if err != nil {
			return
		}
		fmt.Printf("id: [%s]\n", *id)
	}
	{
		inputLable := []string{
			"young", "myope",
		}
		id, err := FindLable(&dbLables, &inputLable)
		if err != nil {
			return
		}
		fmt.Printf("id: [%s]\n", *id)
	}
	{
		inputLable := []string{
			"young", "myope", "no",
		}
		id, err := FindLable(&dbLables, &inputLable)
		if err != nil {
			return
		}
		fmt.Printf("id: [%s]\n", *id)
	}
	{
		inputLable := []string{
			"young", "myope", "no", "reduced",
		}
		id, err := FindLable(&dbLables, &inputLable)
		if err != nil {
			return
		}
		fmt.Printf("id: [%s]\n", *id)
	}
}
