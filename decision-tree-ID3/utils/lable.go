package utils

import (
	"errors"
	"gopkg.in/fatih/set.v0"
)

func FindLable(dbLables *map[string][]string, inputLable *[]string) (*string, error) {
	inputLableSet := set.New(set.ThreadSafe)
	for _, item := range *inputLable {
		inputLableSet.Add(item)
	}
	var id *string = nil
	for key, value := range *dbLables {
		s := set.New(set.ThreadSafe)
		for _, item := range value {
			s.Add(item)
		}
		if len(set.Difference(s, inputLableSet).List()) == 0 &&
			len(set.Difference(inputLableSet, s).List()) == 0 {
			keyTmp := key
			id = &keyTmp
			break
		}
	}
	if id == nil {
		return nil, errors.New("not found")
	}
	return id, nil
}
