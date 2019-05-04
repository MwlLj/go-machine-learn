package handler

import (
	"../buffer"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	modeNormal string = "normal"
	modeAdd    string = "add"
	modeInfer  string = "infer"
)

type CHandler struct {
	data       *buffer.CData
	mode       string
	dataBuffer map[string]string
}

func (this *CHandler) handlerInput(in *string) bool {
	input := strings.TrimSpace(*in)
	buf := bytes.Buffer{}
	after := ""
	for index, item := range input {
		str := string(item)
		buf.WriteString(str)
		if str == " " {
			after = input[index:]
			after = strings.TrimSpace(after)
			break
		}
	}
	var _ = after
	ret := true
	key := buf.String()
	key = strings.TrimSpace(key)
	if this.mode == modeNormal {
		if key == "lables" {
			this.showAllLables()
		} else if key == "quit" || key == "exit" || key == "q" {
			ret = false
		} else if key == "add" {
			this.mode = modeAdd
		} else if key == "infer" {
			this.mode = modeInfer
		}
	} else if this.mode == modeAdd {
		if key == "end" {
			after = strings.TrimSpace(after)
			if after == "" {
				fmt.Println("end之后一定要加分类名称, 如: end 引擎损坏")
				return ret
			}
			this.addData(&after)
			this.mode = modeNormal
			this.dataBuffer = make(map[string]string)
		} else {
			this.addDataBuffer(&input)
		}
	} else if this.mode == modeInfer {
		if key == "end" {
			this.inferResult()
			this.mode = modeNormal
			this.dataBuffer = make(map[string]string)
		} else {
			this.addDataBuffer(&input)
		}
	}
	return ret
}

func (this *CHandler) inferResult() {
	this.data.InferResult(&this.dataBuffer)
}

func (this *CHandler) addData(classify *string) {
	isExist := true
	for k, _ := range this.dataBuffer {
		kTmp := k
		if !this.data.IsExist(&kTmp) {
			isExist = false
			break
		}
	}
	if !isExist {
		fmt.Println("is not exist => update db")
		// update data add field, add new data
		this.data.UpdateData(&this.dataBuffer, classify)
	} else {
		fmt.Println("is exist => add db")
		// add data
		this.data.AddDataToDb(&this.dataBuffer, classify)
	}
	this.data.BuildTree()
}

func (this *CHandler) addDataBuffer(input *string) {
	if strings.Count(*input, ":") > 1 {
		fmt.Println("不能含有多余的冒号")
		return
	}
	vec := strings.Split(*input, ":")
	if len(vec) != 2 {
		fmt.Println("输入的格式不正确, 使用 标签名: 标签值 的格式输入, 如: 是否能够行驶: 是")
		return
	}
	lable := strings.TrimSpace(vec[0])
	lableValue := strings.TrimSpace(vec[1])
	this.dataBuffer[lable] = lableValue
}

func (this *CHandler) showAllLables() {
	lables := this.data.GetLables()
	lablesLen := len(*lables)
	if lablesLen > 0 {
		buf := bytes.Buffer{}
		for index, item := range *lables {
			buf.WriteString(item)
			if index < lablesLen-1 {
				buf.WriteString(", ")
			}
		}
		fmt.Println(buf.String())
	} else {
		fmt.Println("数据库中还没有数据, 请使用 add 命令录入吧")
	}
}

func (this *CHandler) InputHandle() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("[Error] scanf error")
			continue
		}
		input = strings.Replace(input, "\n", "", 1)
		input = strings.Replace(input, "\r", "", 1)
		r := this.handlerInput(&input)
		if !r {
			break
		}
	}
}

func NewHandler(data *buffer.CData) *CHandler {
	handler := CHandler{
		data: data,
		mode: modeNormal,
	}
	handler.dataBuffer = make(map[string]string)
	return &handler
}

func (this *CHandler) PrintMethod() {
	fmt.Println(`
请输入以下选项操作:
	1. 输入 lables
		显示 所欲的标签值
		如:
			lables
	2. 输入 add
		该选项用于录入专家数据, 输入 add 以后, 将进入录入模式, 该模式下, 一次只能录入一条数据
		, 使用 end 来结束, 注意: end 后一定要加 分类名称, 否则将报错
		进入 add 模式后, 就可以添加一条数据, 用下面的方式录入:
		标签名: 就是专家判断的依据
		标签值: 就是依据对应的值
		如:
			是否可以运行: 是
			轮胎是否破损: 否
	3. 输入 infer
		该选项用于推测数据, 输入 infer 以后, 将进入推测模式
		与 add 类似, 使用 end 结束该模式
		进入 infer 模式后, 就可以推测输入的数据, 用下面的方式录入:
		标签名: 就是专家判断的依据
		标签值: 就是依据对应的值
		如:
			是否可以运行: 是
			轮胎是否破损: 否
	4. 输入 quit 或者 exit 或者 q
		退出

	5. 完整示例:
		(1) 添加:

		add
		是否有房: 是
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 否
		学历是否研究生及以上: 否
		end 不见面

		add
		是否有房: 否
		是否有车: 否
		是否有存款: 否
		身高是否大于180: 是
		学历是否研究生及以上: 是
		end 可能见面

		add
		是否有房: 是
		是否有车: 是
		是否有存款: 否
		身高是否大于180: 否
		学历是否研究生及以上: 否
		end 不见面

		add
		是否有房: 是
		是否有车: 否
		是否有存款: 否
		身高是否大于180: 是
		学历是否研究生及以上: 是
		end 见面

		add
		是否有房: 是
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 是
		学历是否研究生及以上: 是
		end 见面

		add
		是否有房: 是
		是否有车: 是
		是否有存款: 否
		身高是否大于180: 否
		学历是否研究生及以上: 否
		end 不见面

		add
		是否有房: 否
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 是
		学历是否研究生及以上: 是
		end 不见面

		add
		是否有房: 否
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 否
		学历是否研究生及以上: 否
		end 不见面

		add
		是否有房: 否
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 是
		学历是否研究生及以上: 否
		end 不见面

		add
		是否有房: 否
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 是
		学历是否研究生及以上: 是
		end 不见面

		(2) 推断:
		infer
		是否有房: 否
		是否有车: 是
		是否有存款: 是
		身高是否大于180: 是
		学历是否研究生及以上: 否
		end
		`)
}
