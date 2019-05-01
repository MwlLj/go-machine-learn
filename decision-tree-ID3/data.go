package tree

func SplitDataSet(dataSet *[][]string, featureIndex int, featureValue *string) *[][]string {
	var retVec [][]string
	for _, item := range *dataSet {
		value := item[featureIndex]
		if value == *featureValue {
			itemLen := len(item)
			itemCopy := make([]string, itemLen, itemLen)
			copy(itemCopy, item)
			vec := itemCopy[:featureIndex]
			vec = append(vec, itemCopy[featureIndex+1:]...)
			retVec = append(retVec, vec)
		}
	}
	return &retVec
}
