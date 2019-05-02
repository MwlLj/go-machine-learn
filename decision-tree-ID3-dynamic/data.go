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

// find greater than featureIndex sub set
func FindGreaterThanFeatureIndexSet(dataSet *[][]string, featureIndex int) *[][]string {
	var retVec [][]string
	for _, item := range *dataSet {
		featureLen := len(item) - 1
		if featureLen >= featureIndex+1 {
			retVec = append(retVec, item)
		}
	}
	return &retVec
}

// split by featureNumber
type SubDataSet struct {
	FeatureNumberMin int
	FeatureNumberMax int
	DataSet          *[][]string
}

func SplitDataSetByFeatureNumber(dataSet *[][]string) *[]SubDataSet {
	mapTmp := map[int]SubDataSet{}
	for _, item := range *dataSet {
		number := len(item)
		if _, ok := mapTmp[number]; !ok {
			set := [][]string{}
			set = append(set, item)
			data := SubDataSet{
				FeatureNumberMax: number,
				DataSet:          &set,
			}
			mapTmp[number] = data
		} else {
			*mapTmp[number].DataSet = append(*mapTmp[number].DataSet, item)
		}
	}
	retVec := []SubDataSet{}
	for _, item := range mapTmp {
		retVec = append(retVec, item)
	}
	return &retVec
}
