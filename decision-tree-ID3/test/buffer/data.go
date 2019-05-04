package buffer

import (
	db "../infer_db"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-machine-learn/decision-tree-ID3"
	"github.com/MwlLj/gotools/timetool"
	"github.com/satori/go.uuid"
	"log"
)

var _ = fmt.Println

type CData struct {
	dbHandler db.CDbHandler
	lables    []string
	treeData  *tree.Node
}

func (this *CData) UpdateLables(lables *[]string) {
	this.lables = this.lables[0:0]
	for _, item := range *lables {
		this.lables = append(this.lables, item)
	}
}

func (this *CData) UpdateLablesByMap(lables *map[string]string) {
	this.lables = this.lables[0:0]
	for key, _ := range *lables {
		this.lables = append(this.lables, key)
	}
	fmt.Println(this.lables)
}

func (this *CData) GetLables() *[]string {
	retLables := []string{}
	for _, item := range this.lables {
		retLables = append(retLables, item)
	}
	return &retLables
}

func (this *CData) GetLabelsLen() int {
	return len(this.lables)
}

func (this *CData) IsExist(lable *string) bool {
	for _, item := range this.lables {
		if *lable == item {
			return true
		}
	}
	return false
}

func (this *CData) BuildTree() {
	fmt.Println("-------开始构建树-------")

	dataSet := [][]string{}
	// get all data
	output := []db.CGetAllDataOutput{}
	err, _ := this.dbHandler.GetAllData(&output)
	if err != nil {
		log.Printf("getAllData error, err: %v\n", err)
		return
	}
	for _, item := range output {
		lableKv := map[string]string{}
		err = json.Unmarshal([]byte(item.LableKV), &lableKv)
		if err != nil {
			log.Println("decode lableKv error")
			continue
		}
		vec := []string{}
		for _, lable := range this.lables {
			vec = append(vec, lableKv[lable])
		}
		vec = append(vec, item.Classify)
		dataSet = append(dataSet, vec)
	}
	node := tree.CreateTree(&dataSet, &this.lables)
	b, err := json.Marshal(node)
	if err != nil {
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return
	}
	fmt.Println("-------树构建完成-------")
	fmt.Println("-----------------------")
	fmt.Println("树结构:")
	fmt.Println(out.String())
	fmt.Println("-----------------------")
	// write tree to db
	if this.writeTreeToTree(node) != nil {
		return
	}
	this.updateTree()
}

func (this *CData) writeTreeToTree(node *tree.Node) error {
	b, err := json.Marshal(node)
	if err != nil {
		log.Printf("encode tree json error, err: %v\n", err)
		return err
	}
	input := db.CUpdateTreeInput{}
	input.Tree = string(b)
	err, _ = this.dbHandler.UpdateTree(&input)
	if err != nil {
		log.Printf("addTree to db error, err: %v\n", err)
		return err
	}
	return nil
}

func (this *CData) InferResult(kv *map[string]string) {
	values := []string{}
	lables := []string{}
	for key, value := range *kv {
		values = append(values, value)
		lables = append(lables, key)
	}
	classify := tree.FindByOrderFeature(this.treeData, &values, &lables)
	if classify != nil {
		fmt.Printf("推测结果: %s\n", *classify)
	} else {
		fmt.Println("无法推测")
	}
}

func (this *CData) UpdateData(kv *map[string]string, classify *string) {
	// get all data
	output := []db.CGetAllDataOutput{}
	err, _ := this.dbHandler.GetAllData(&output)
	if err != nil {
		log.Printf("getAllData error, err: %v\n", err)
		return
	}
	input := []db.CUpdateThenAddDataInput{}
	for _, item := range output {
		lableKv := map[string]string{}
		err = json.Unmarshal([]byte(item.LableKV), &lableKv)
		if err != nil {
			log.Println("decode lableKv error")
			continue
		}
		this.updateLableKv(&lableKv, kv)
		// encode
		b, err := json.Marshal(&lableKv)
		if err != nil {
			log.Println("encode lableKv error")
			continue
		}
		// update
		in := db.CUpdateThenAddDataInput{}
		in.LableKV = string(b)
		in.Classify = item.Classify
		in.DataUuid = item.DataUuid
		input = append(input, in)
	}
	addInput := this.joinAddInput(kv, classify)
	err, _ = this.dbHandler.UpdateThenAddData(&input, addInput)
	if err != nil {
		log.Printf("updateData error, err: %v\n", err)
		return
	}
	this.updateLables()
}

func (this *CData) joinAddInput(kv *map[string]string, classify *string) *[]db.CAddDataInput {
	b, err := json.Marshal(kv)
	if err != nil {
		log.Printf("encode data error, err: %v\n", err)
		return nil
	}
	uid, err := this.genUuid()
	if err != nil {
		return nil
	}
	input := []db.CAddDataInput{}
	now := timetool.GetNowSecondFormat()
	in := db.CAddDataInput{}
	in.DataUuid = *uid
	in.LableKV = string(b)
	in.Classify = *classify
	in.CreateTime = now
	in.UpdateTime = now
	input = append(input, in)
	return &input
}

func (this *CData) AddDataToDb(kv *map[string]string, classify *string) {
	input := this.joinAddInput(kv, classify)
	err, _ := this.dbHandler.AddData(input)
	if err != nil {
		log.Printf("addData error, err: %v\n", err)
		return
	}
}

func (this *CData) updateLableKv(oldKv *map[string]string, newKv *map[string]string) {
	for key, _ := range *newKv {
		if _, ok := (*oldKv)[key]; !ok {
			// if not exist, set value is ""
			(*oldKv)[key] = ""
		}
	}
}

func (this *CData) updateTree() {
	output := db.CGetTreeOutput{}
	err, _ := this.dbHandler.GetTree(&output)
	if err != nil {
		log.Printf("getTree error from db, err: %v\n", err)
		return
	}
	node := tree.Node{}
	err = json.Unmarshal([]byte(output.Tree), &node)
	if err != nil {
		log.Printf("decode tree json error, err: %v\n", err)
		return
	}
	this.treeData = &node
}

func (this *CData) updateLables() {
	output := db.CGetOneDataOutput{}
	err, count := this.dbHandler.GetOneData(&output)
	if err != nil {
		log.Fatalln("get data error")
	}
	if count == 0 {
		return
	}
	lableKv := map[string]string{}
	err = json.Unmarshal([]byte(output.LableKV), &lableKv)
	if err != nil {
		log.Fatalln("decode lable kv error")
	}
	this.UpdateLablesByMap(&lableKv)
}

func NewData(dbHandler db.CDbHandler) *CData {
	data := CData{
		dbHandler: dbHandler,
	}
	data.updateLables()
	data.updateTree()
	return &data
}

func (this *CData) genUuid() (*string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Printf("general uuid error, err: %v", err)
		return nil, err
	}
	id := uid.String()
	return &id, nil
}
