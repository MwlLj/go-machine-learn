package infer_db

type CAddDataInput struct {
	DataUuid string
	DataUuidIsValid bool
	LableKV string
	LableKVIsValid bool
	Classify string
	ClassifyIsValid bool
	CreateTime string
	CreateTimeIsValid bool
	UpdateTime string
	UpdateTimeIsValid bool
}

type CUpdateDataInput struct {
	DataUuid string
	DataUuidIsValid bool
	LableKV string
	LableKVIsValid bool
	Classify string
	ClassifyIsValid bool
	UpdateTime string
	UpdateTimeIsValid bool
}

type CUpdateThenAddDataInput struct {
	DataUuid string
	DataUuidIsValid bool
	LableKV string
	LableKVIsValid bool
	Classify string
	ClassifyIsValid bool
	UpdateTime string
	UpdateTimeIsValid bool
}

type CGetAllDataOutput struct {
	DataUuid string
	DataUuidIsValid bool
	LableKV string
	LableKVIsValid bool
	Classify string
	ClassifyIsValid bool
}

type CGetOneDataOutput struct {
	LableKV string
	LableKVIsValid bool
	Classify string
	ClassifyIsValid bool
}

type CUpdateTreeInput struct {
	Tree string
	TreeIsValid bool
}

type CGetTreeOutput struct {
	Tree string
	TreeIsValid bool
}

