// 用于将Instruct指令序列化为字节码文件

package ss

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// 用于将Instruct指令序列化为字节码文件

// CodeCompiler 代码编解码接口
type CodeCompiler interface {
	UnmarshalBin([]byte) (any, error)
	MarshalBin() ([]byte, error)
}

// Bin 字节码文件的结构体
type Bin struct {
	head    string
	version []uint8
	files   []string
	code    []Instruct
}

// Uint64ToByte int转字节数组
func Uint64ToByte(i uint64) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return b.Bytes()
}

// ByteToUint64 字节数组转int
func ByteToUint64(b []byte) uint64 {
	var i uint64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return i
}

// Uint16ToByte uint16转字节数组
func Uint16ToByte(i uint16) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return b.Bytes()
}

// ByteToUint16 字节数组转uint16
func ByteToUint16(b []byte) uint16 {
	var i uint16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return i
}

// Uint32ToByte uint32转字节数组
func Uint32ToByte(i uint32) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return b.Bytes()
}

// ByteToUint32 字节数组转uint32
func ByteToUint32(b []byte) uint32 {
	var i uint32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return i
}

// IntToByte int转字节数组
func IntToByte[T uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64](i T) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return b.Bytes()
}

// ByteToInt 字节数组转uint
func ByteToInt[T uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64](b []byte) T {
	var i T
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return i
}

// Uint8ToByte uint8转字节数组
func Uint8ToByte(i uint8) []byte {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return b.Bytes()
}

// ByteToUint8 字节数组转uint8
func ByteToUint8(b []byte) uint8 {
	var i uint8
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		panic(NewError("Error", err.Error()))
	}
	return i
}

// 字节码编译器,将指令序列化成字节码文件
func build(c []Instruct) []byte {
	data := headToBin()
	data = append(data, filesToBin()...)
	bin, _ := MarshalBinArray(c)
	data = append(data, bin...)
	return data
}

// 将字节码文件反序列化为Instruct指令序列
func unBuild(bin []byte) *Bin {
	b := Bin{}
	s, v := binToHead(bin)
	f, p := binToFiles(bin[7:])
	c, _ := UnmarshalBinArray(bin[p+7:])
	b.version = v
	b.files = f
	b.head = s
	b.code = c
	return &b
}

// 返回字节码头
func headToBin() []byte {
	data := make([]byte, 7)
	//写入魔数,JWJS的ASCII码
	data[0] = 74 //J
	data[1] = 87 //W
	data[2] = 74 //J
	data[3] = 83 //S
	//写入版本号
	copy(data[4:7], version)
	return data
}

// 将字节码头不进行反序列化
func binToHead(c []byte) (string, []uint8) {
	if len(c) < 7 {
		return "", nil
	}
	h := string(c[0]) + string(c[1]) + string(c[2]) + string(c[3])
	if h != "JWJS" {
		panic("文件类型校验失败")
	}
	return h, c[4:7]
}

// filesToBin 文件列表
/*
   文件数量(2字节),[文件名长度,文件名内容]n个
*/
func filesToBin() []byte {
	//文件列表,数组长度2个字节
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(len(getGlobalFiles())))
	for _, v := range getGlobalFiles() {
		name := []byte(v)
		length := make([]byte, 2)
		binary.BigEndian.PutUint16(length, uint16(len(name)))
		data = append(data, length...)
		data = append(data, name...)
	}
	return data
}

// 将字节转为文件列表
func binToFiles(c []byte) ([]string, int) {
	length := binary.BigEndian.Uint16(c[:2])
	files := make([]string, length)
	p := 2
	for i := range length {
		l := int(binary.BigEndian.Uint16(c[p : p+2]))
		p = p + 2
		name := string(c[p : p+l])
		p = p + l
		files[i] = name
	}
	return files, p
}

// UnmarshalBinArray 将字节码反序列化成指令
func UnmarshalBinArray(bytes []byte) ([]Instruct, error) {
	if len(bytes) == 0 {
		return make([]Instruct, 0), nil
	}
	size := uint64(len(bytes))
	data := make([]Instruct, 0)
	head := bytes[0]
	length := ByteToUint64(bytes[1:9])
	start := uint64(9)
	last := length + start
	for {
		var d any
		switch head {
		case 20:
			c := Operate(0)
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 21:
			c := command(0)
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 22:
			c := block{}
			d, _ = c.UnmarshalBin(bytes[start:last])
			d = block(d.([]Instruct))
		case 23:
			c := iteration{}
			d, _ = c.UnmarshalBin(bytes[start:last])
			d = iteration(d.([]Instruct))
		case 24:
			c := createObject{}
			d, _ = c.UnmarshalBin(bytes[start:last])
			d = createObject(d.([]Instruct))
		case 25:
			c := to(0)
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 26:
			c := array(0)
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 27:
			c := loadVal{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 28:
			c := loadVar{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 29:
			c := cmd{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 30:
			c := modifyVar{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 31:
			c := createVar{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 32:
			c := modifyProperty{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 33:
			c := property{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 34:
			c := call{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 35:
			c := importModule{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 36:
			c := exportItem{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 37:
			c := function{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 38:
			c := try{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 39:
			c := skip(0)
			d, _ = c.UnmarshalBin(bytes[start:last])
		case 40:
			c := forOf{}
			d, _ = c.UnmarshalBin(bytes[start:last])
		default:
			panic(fmt.Sprint("指令类型错误:", head))
		}
		data = append(data, d.(Instruct))
		if last >= size {
			break
		}
		head = bytes[last]
		start = last + 9
		length = ByteToUint64(bytes[last+1 : last+9])
		last = length + start
	}
	return data, nil
}

// MarshalBinArray 将指令序列化成字节码
func MarshalBinArray(codes []Instruct) ([]byte, error) {
	data := make([]byte, 0)
	for _, v := range codes {
		d, _ := v.MarshalBin()
		var head byte = 20
		switch v.(type) {
		case Operate:
			head = 20
		case command:
			head = 21
		case block:
			head = 22
		case iteration:
			head = 23
		case createObject:
			head = 24
		case to:
			head = 25
		case array:
			head = 26
		case loadVal:
			head = 27
		case loadVar:
			head = 28
		case cmd:
			head = 29
		case modifyVar:
			head = 30
		case createVar:
			head = 31
		case modifyProperty:
			head = 32
		case property:
			head = 33
		case call:
			head = 34
		case importModule:
			head = 35
		case exportItem:
			head = 36
		case function:
			head = 37
		case try:
			head = 38
		case skip:
			head = 39
		case forOf:
			head = 40
		default:
			panic("类型错误")
		}
		data = append(data, []byte{head}...)
		data = append(data, Uint64ToByte(uint64(len(d)))...)
		data = append(data, d...)
	}
	return data, nil
}

func (t block) UnmarshalBin(bytes []byte) (any, error) {
	return UnmarshalBinArray(bytes)
}

func (t block) MarshalBin() ([]byte, error) {
	return MarshalBinArray(t)
}

func (t iteration) UnmarshalBin(bytes []byte) (any, error) {
	return UnmarshalBinArray(bytes)
}

func (t iteration) MarshalBin() ([]byte, error) {
	return MarshalBinArray(t)
}

func (t array) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint32(bytes[0:4])
	return array(v1), nil
}

func (t array) MarshalBin() ([]byte, error) {
	v1 := Uint32ToByte(uint32(t))
	return v1, nil
}
func (t to) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToInt[int32](bytes[0:4])
	return to(v1), nil
}

func (t to) MarshalBin() ([]byte, error) {
	v1 := IntToByte(int32(t))
	return v1, nil
}

func (c call) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint8(bytes[0:1])
	v2 := ByteToUint8(bytes[1:2])
	v3 := ByteToUint8(bytes[2:3])
	v4 := ByteToUint8(bytes[3:4])
	v5 := ByteToUint16(bytes[4:6])
	vv1 := bytes[6 : 6+v1]
	vv2 := bytes[6+v1 : 6+v1+v2]
	vv3 := bytes[6+v1+v2 : 6+v1+v2+v3]
	vv4 := bytes[6+v1+v2+v3 : 6+v1+v2+v3+v4]
	vv5 := bytes[6+v1+v2+v3+v4 : 6+uint16(v1)+uint16(v2)+uint16(v3)+uint16(v4)+v5]
	c.index = int(ByteToUint8(vv1))
	c.state = int(ByteToUint8(vv2))
	if int(ByteToUint8(vv3)) == 1 {
		c.constructor = true
	}
	if int(ByteToUint8(vv4)) == 1 {
		c.optional = true
	}
	var src *Src
	if len(vv5) > 0 {
		src = &Src{}
		src.UnmarshalBin(vv5)
	}
	c.src = src
	return c, nil
}

func (c call) MarshalBin() ([]byte, error) {
	v1 := Uint8ToByte(uint8(c.index))
	v2 := Uint8ToByte(uint8(c.state))
	var v3 []byte
	if c.constructor {
		v3 = Uint8ToByte(uint8(1))
	} else {
		v3 = Uint8ToByte(uint8(0))
	}
	var v4 []byte
	if c.optional {
		v4 = Uint8ToByte(uint8(1))
	} else {
		v4 = Uint8ToByte(uint8(0))
	}

	v5 := make([]byte, 0)
	if c.src != nil {
		v5, _ = c.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint8ToByte(uint8(len(v1)))...)
	data = append(data, Uint8ToByte(uint8(len(v2)))...)
	data = append(data, Uint8ToByte(uint8(len(v3)))...)
	data = append(data, Uint8ToByte(uint8(len(v4)))...)
	data = append(data, Uint16ToByte(uint16(len(v5)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	data = append(data, v3...)
	data = append(data, v4...)
	data = append(data, v5...)
	return data, nil
}

func (l property) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint8(bytes[2:3])
	v3 := ByteToUint8(bytes[3:4])
	v4 := ByteToUint16(bytes[4:6])
	vv1 := bytes[6 : 6+v1]
	vv2 := bytes[6+v1 : 6+v1+uint16(v2)]
	vv3 := bytes[6+v1+uint16(v2) : 6+v1+uint16(v2)+uint16(v3)]
	vv4 := bytes[6+v1+uint16(v2)+uint16(v3) : 6+v1+uint16(v2)+uint16(v3)+v4]
	l.name = string(vv1)
	if vv2[0] == 1 {
		l.delete = true
	} else {
		l.delete = false
	}
	if vv3[0] == 1 {
		l.optional = true
	} else {
		l.optional = false
	}
	var src *Src
	if len(vv4) > 0 {
		src = &Src{}
		src.UnmarshalBin(vv4)
	}
	l.src = src
	return l, nil
}

func (l property) MarshalBin() ([]byte, error) {
	v1 := []byte(l.name)
	var v2 []byte
	if l.delete {
		v2 = []byte{1}
	} else {
		v2 = []byte{0}
	}

	var v3 []byte
	if l.optional {
		v3 = []byte{1}
	} else {
		v3 = []byte{0}
	}

	v4, _ := l.src.MarshalBin()
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint8ToByte(uint8(len(v2)))...)
	data = append(data, Uint8ToByte(uint8(len(v3)))...)
	data = append(data, Uint16ToByte(uint16(len(v4)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	data = append(data, v3...)
	data = append(data, v4...)
	return data, nil
}

func (m modifyProperty) UnmarshalBin(bytes []byte) (any, error) {
	vl := ByteToUint16(bytes[0:2])
	sl := ByteToUint16(bytes[2:4])
	vv := bytes[4 : 4+vl]
	sv := bytes[4+vl : 4+vl+sl]
	m.name = string(vv)
	var src *Src
	if len(sv) > 0 {
		src = &Src{}
		src.UnmarshalBin(sv)
	}
	m.src = src
	return m, nil
}

func (m modifyProperty) MarshalBin() ([]byte, error) {
	v1 := []byte(m.name)
	v2 := make([]byte, 0)
	if m.src != nil {
		v2, _ = m.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint16ToByte(uint16(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}

func (c createVar) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint8(bytes[2:3])
	v3 := ByteToUint16(bytes[3:5])
	vv1 := bytes[5 : 5+v1]
	vv2 := bytes[5+v1 : 5+v1+uint16(v2)]
	vv3 := bytes[5+v1+uint16(v2) : 5+v1+uint16(v2)+v3]
	c.name = string(vv1)
	c.mod = modifier(ByteToUint8(vv2))
	var src *Src
	if len(vv3) > 0 {
		src = &Src{}
		src.UnmarshalBin(vv3)
	}
	c.src = src
	return c, nil
}

func (c createVar) MarshalBin() ([]byte, error) {
	v1 := []byte(c.name)
	v2 := Uint8ToByte(uint8(c.mod))
	v3 := make([]byte, 0)
	if c.src != nil {
		v3, _ = c.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint8ToByte(uint8(len(v2)))...)
	data = append(data, Uint16ToByte(uint16(len(v3)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	data = append(data, v3...)
	return data, nil
}

func (m modifyVar) UnmarshalBin(bytes []byte) (any, error) {
	vl := ByteToUint16(bytes[0:2])
	sl := ByteToUint16(bytes[2:4])
	vv := bytes[4 : 4+vl]
	sv := bytes[4+vl : 4+vl+sl]
	m.name = string(vv)
	var src *Src
	if len(sv) > 0 {
		src = &Src{}
		src.UnmarshalBin(sv)
	}
	m.src = src
	return m, nil
}

func (m modifyVar) MarshalBin() ([]byte, error) {
	v1 := []byte(m.name)
	v2 := make([]byte, 0)
	if m.src != nil {
		v2, _ = m.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint16ToByte(uint16(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}

func (c cmd) UnmarshalBin(bytes []byte) (any, error) {
	vl := ByteToUint8(bytes[0:1])
	sl := ByteToUint16(bytes[1:3])
	vv := bytes[3 : 3+vl]
	sv := bytes[3+vl : 3+uint16(vl)+sl]
	c.c = command(ByteToUint8(vv))
	var src *Src
	if len(sv) > 0 {
		src = &Src{}
		src.UnmarshalBin(sv)
	}
	c.src = src
	return c, nil
}

func (c cmd) MarshalBin() ([]byte, error) {
	v1 := Uint8ToByte(uint8(c.c))
	v2 := make([]byte, 0)
	if c.src != nil {
		v2, _ = c.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint8ToByte(uint8(len(v1)))...)
	data = append(data, Uint16ToByte(uint16(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}

func (l loadVar) UnmarshalBin(bytes []byte) (any, error) {
	vl := ByteToUint16(bytes[0:2])
	sl := ByteToUint16(bytes[2:4])
	vv := bytes[4 : 4+vl]
	sv := bytes[4+vl : 4+vl+sl]
	l.name = string(vv)
	var src *Src
	if len(sv) > 0 {
		src = &Src{}
		src.UnmarshalBin(sv)
	}
	l.src = src
	return l, nil
}

func (l loadVar) MarshalBin() ([]byte, error) {
	v1 := []byte(l.name)
	v2 := make([]byte, 0)
	if l.src != nil {
		v2, _ = l.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint16ToByte(uint16(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}

func (l loadVal) UnmarshalBin(bytes []byte) (any, error) {
	//1:bigint,2:boolean,3:null,4:undefined,5:number,6:regexp,7:string
	vl := ByteToUint16(bytes[0:2])
	sl := ByteToUint16(bytes[2:4])
	vv := bytes[4 : 4+vl]
	sv := bytes[4+vl : 4+vl+sl]
	var value JavaScript
	switch vv[0] {
	case 1:
		value = GoToBigInt(0)
	case 2:
		value = GoToBoolean(false)
	case 3:
		value = NewNull()
	case 4:
		value = NewUndefined()
	case 5:
		value = GoToNumber(0)
	case 6:
		value = GoToRegExp(`/^\d+$/`)
	case 7:
		value = GoToString("")
	case 8:
		value = GoToBuffer([]byte{})
	default:
		value = NewUndefined()
	}
	if c, ok := value.(CodeCompiler); ok {
		c.UnmarshalBin(vv[1:])
		var src *Src
		if len(sv) > 0 {
			src = &Src{}
			src.UnmarshalBin(sv)
		}
		l.value = value
		l.src = src
		return l, nil
	} else {
		panic("未实现CodeCompiler接口")
	}
}

func (l loadVal) MarshalBin() ([]byte, error) {
	var v1 []byte
	if c, ok := l.value.(CodeCompiler); ok {
		v1, _ = c.MarshalBin()
	} else {
		panic("未实现CodeCompiler接口")
	}
	v2 := make([]byte, 0)
	if l.src != nil {
		v2, _ = l.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint16ToByte(uint16(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}

func (c createObject) UnmarshalBin(bytes []byte) (any, error) {
	return UnmarshalBinArray(bytes)
}

func (c createObject) MarshalBin() ([]byte, error) {
	return MarshalBinArray(c)
}
func (c command) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint8(bytes[0:1])
	t1 := command(v1)
	return t1, nil
}

func (c command) MarshalBin() ([]byte, error) {
	v1 := Uint8ToByte(uint8(c))
	return v1, nil
}

func (o Operate) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint8(bytes[0:1])
	t1 := Operate(v1)
	return t1, nil
}

func (o Operate) MarshalBin() ([]byte, error) {
	v1 := Uint8ToByte(uint8(o))
	return v1, nil
}

func (t function) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint64(bytes[2:10])
	vv1 := bytes[10 : 10+v1]
	vv2 := bytes[10+v1 : 10+uint64(v1)+(v2)]
	formal := make([]string, 0)
	json.Unmarshal(vv1, &formal)
	t.formal = formal
	t.body, _ = UnmarshalBinArray(vv2)
	return t, nil
}

func (t function) MarshalBin() ([]byte, error) {
	v1, _ := json.Marshal(t.formal)
	v2, _ := MarshalBinArray(t.body)
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint64ToByte(uint64(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}
func (c catch) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint64(bytes[2:10])
	vv1 := bytes[10 : 10+v1]
	vv2 := bytes[10+v1 : 10+uint64(v1)+v2]
	c.formal = string(vv1)
	c.body, _ = UnmarshalBinArray(vv2)
	return c, nil
}

func (c catch) MarshalBin() ([]byte, error) {
	v1 := []byte(c.formal)
	v2, _ := MarshalBinArray(c.body)
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint64ToByte(uint64(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}

func (t try) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint64(bytes[0:8])
	v2 := ByteToUint64(bytes[8:16])
	v3 := ByteToUint64(bytes[16:24])
	vv1 := bytes[24 : 24+v1]
	vv2 := bytes[24+v1 : 24+v1+v2]
	vv3 := bytes[24+v1+v2 : 24+v1+v2+v3]
	t.try, _ = UnmarshalBinArray(vv1)
	t.finally, _ = UnmarshalBinArray(vv2)
	cat := catch{}
	bin, _ := cat.UnmarshalBin(vv3)
	t.catch = bin.(catch)
	return t, nil
}

func (t try) MarshalBin() ([]byte, error) {
	v1, _ := MarshalBinArray(t.try)
	v2, _ := MarshalBinArray(t.finally)
	v3, _ := t.catch.MarshalBin()
	data := make([]byte, 0)
	data = append(data, Uint64ToByte(uint64(len(v1)))...)
	data = append(data, Uint64ToByte(uint64(len(v2)))...)
	data = append(data, Uint64ToByte(uint64(len(v3)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	data = append(data, v3...)
	return data, nil
}

func (t exportItem) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint16(bytes[2:4])
	v3 := ByteToUint16(bytes[4:6])
	vv1 := bytes[6 : 6+v1]
	vv2 := bytes[6+v1 : 6+v1+v2]
	vv3 := bytes[6+v1+v2 : 6+v1+v2+v3]
	t.name = string(vv1)
	t.as = string(vv2)
	var src *Src
	if len(vv3) > 0 {
		src = &Src{}
		src.UnmarshalBin(vv3)
	}
	t.src = src
	return t, nil
}

func (t exportItem) MarshalBin() ([]byte, error) {
	v1 := []byte(t.name)
	v2 := []byte(t.as)
	v3 := make([]byte, 0)
	if t.src != nil {
		v3, _ = t.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint16ToByte(uint16(len(v2)))...)
	data = append(data, Uint16ToByte(uint16(len(v3)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	data = append(data, v3...)
	return data, nil
}

func (m importModule) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint64(bytes[2:10])
	v3 := ByteToUint16(bytes[10:12])
	vv1 := bytes[12 : 12+v1]
	vv2 := bytes[12+v1 : 12+uint64(v1)+v2]
	vv3 := bytes[12+uint64(v1)+v2 : 12+uint64(v1)+v2+uint64(v3)]
	im := &importModule{}
	json.Unmarshal(vv1, im)
	im.code, _ = UnmarshalBinArray(vv2)
	var src *Src
	if len(vv3) > 0 {
		src = &Src{}
		src.UnmarshalBin(vv3)
	}
	im.src = src
	return *im, nil
}

func (m importModule) MarshalBin() ([]byte, error) {
	v1, _ := json.Marshal(m)
	v2, _ := MarshalBinArray(m.code)
	v3 := make([]byte, 0)
	if m.src != nil {
		v3, _ = m.src.MarshalBin()
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint64ToByte(uint64(len(v2)))...)
	data = append(data, Uint16ToByte(uint16(len(v3)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	data = append(data, v3...)
	return data, nil
}

func (s *Src) UnmarshalBin(bytes []byte) (any, error) {
	startLine := ByteToUint16(bytes[0:2])
	startColumn := ByteToUint16(bytes[2:4])
	file := ByteToUint16(bytes[4:6])
	s.StartLine = int(startLine)
	s.StartColumn = int(startColumn)
	s.File = int(file)
	return nil, nil
}

func (s *Src) MarshalBin() ([]byte, error) {
	startLine := Uint16ToByte(uint16(s.StartLine))
	startColumn := Uint16ToByte(uint16(s.StartColumn))
	file := Uint16ToByte(uint16(s.File))
	return append(append(startLine, startColumn...), file...), nil
}

func (t skip) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToInt[uint32](bytes[0:4])
	return skip(v1), nil
}

func (t skip) MarshalBin() ([]byte, error) {
	v1 := IntToByte[uint32](uint32(t))
	return v1, nil
}

func (f forOf) UnmarshalBin(bytes []byte) (any, error) {
	v1 := ByteToUint16(bytes[0:2])
	v2 := ByteToUint8(bytes[2:3])
	vv1 := bytes[3 : 3+v1]
	vv2 := bytes[3+v1 : 3+v1+uint16(v2)]
	f.vars = make([]string, 0)
	json.Unmarshal(vv1, &f.vars)
	if int(ByteToUint8(vv2)) == 1 {
		f.list = true
	}
	return f, nil
}

func (f forOf) MarshalBin() ([]byte, error) {
	v1, _ := json.Marshal(f.vars)
	var v2 []byte
	if f.list {
		v2 = Uint8ToByte(uint8(1))
	} else {
		v2 = Uint8ToByte(uint8(0))
	}
	data := make([]byte, 0)
	data = append(data, Uint16ToByte(uint16(len(v1)))...)
	data = append(data, Uint8ToByte(uint8(len(v2)))...)
	data = append(data, v1...)
	data = append(data, v2...)
	return data, nil
}
