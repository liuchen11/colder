package parser

/***************************************************************************************
**文件名：NameTable.go
**包名称：parser
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"fmt"
)

//符号表类型
type NameTable struct {
	//类型变量
	TypeMap map[string]string
	//单变量
	StringMap map[string]string
	IntMap    map[string]int
	BoolMap   map[string]bool
	//数组变量
	StringArrMap map[string][]string
	IntArrMap    map[string][]int
	BoolArrMap   map[string][]bool
}

//构造函数
func NewNameTable() *NameTable {
	ret := new(NameTable)
	ret.TypeMap = make(map[string]string)
	ret.StringMap = make(map[string]string)
	ret.IntMap = make(map[string]int)
	ret.BoolMap = make(map[string]bool)
	ret.StringArrMap = make(map[string][]string)
	ret.IntArrMap = make(map[string][]int)
	ret.BoolArrMap = make(map[string][]bool)
	return ret
}

func CopyNameTable(a *NameTable) *NameTable {
	ret := new(NameTable)
	ret.TypeMap = make(map[string]string)
	ret.StringMap = make(map[string]string)
	ret.IntMap = make(map[string]int)
	ret.BoolMap = make(map[string]bool)
	ret.StringArrMap = make(map[string][]string)
	ret.IntArrMap = make(map[string][]int)
	ret.BoolArrMap = make(map[string][]bool)

	for k, v := range a.TypeMap {
		ret.TypeMap[k] = v
	}
	for k, v := range a.IntMap {
		ret.IntMap[k] = v
	}
	for k, v := range a.BoolMap {
		ret.BoolMap[k] = v
	}
	for k, v := range a.StringMap {
		ret.StringMap[k] = v
	}
	for k, v := range a.IntArrMap {
		size := len(v)
		ret.IntArrMap[k] = make([]int, size, size)
		for i := 1; i < size; i++ {
			ret.IntArrMap[k][i] = v[i]
		}
	}
	for k, v := range a.BoolArrMap {
		size := len(v)
		ret.BoolArrMap[k] = make([]bool, size, size)
		for i := 1; i < size; i++ {
			ret.BoolArrMap[k][i] = v[i]
		}
	}
	for k, v := range a.StringArrMap {
		size := len(v)
		ret.StringArrMap[k] = make([]string, size, size)
		for i := 1; i < size; i++ {
			ret.StringArrMap[k][i] = v[i]
		}
	}
	return ret
}

//添加变量
func (a *NameTable) AddVariable(v_type string, v_name string) bool {
	_, ok := a.TypeMap[v_name]
	if ok == true {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	} else {
		switch v_type {
		case "string":
			a.StringMap[v_name] = ""
		case "int":
			a.IntMap[v_name] = 0
		case "bool":
			a.BoolMap[v_name] = false
		default:
			fmt.Println("NameTableError : Cannot Find the type of " + v_type)
			return false
		}
		a.TypeMap[v_name] = v_type
		return true
	}
}

//添加数组
func (a *NameTable) AddArray(v_type string, v_name string, length int) bool {
	_, ok := a.TypeMap[v_name]
	if ok == true {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	} else {
		switch v_type {
		case "stringarr":
			a.StringArrMap[v_name] = make([]string, length, length)
		case "intarr":
			a.IntArrMap[v_name] = make([]int, length, length)
		case "boolarr":
			fmt.Println(v_name)
			a.BoolArrMap[v_name] = make([]bool, length, length)
		default:
			fmt.Println("NameTableError : Cannot Find the type of " + v_type)
			return false
		}
		a.TypeMap[v_name] = v_type
		return true
	}
}

//获得变量的类型
func (a *NameTable) GetType(v_name string) string {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return "NULL"
	}
	return v_type
}

//从符号表中获得相应变量的值
func (a *NameTable) GetString(v_name string) string {
	return a.StringMap[v_name]
}
func (a *NameTable) GetInt(v_name string) int {
	return a.IntMap[v_name]
}
func (a *NameTable) GetBool(v_name string) bool {
	return a.BoolMap[v_name]
}

//获得数组长度
func (a *NameTable) GetIntlen(v_name string) int {
	return len(a.IntArrMap[v_name])
}
func (a *NameTable) GetStringlen(v_name string) int {
	return len(a.StringArrMap[v_name])
}
func (a *NameTable) GetBoollen(v_name string) int {
	return len(a.BoolArrMap[v_name])
}

//从符号表中获得相应数组的对应位置的值
func (a *NameTable) GetStringArr(v_name string, label int) (string, bool) {
	if len(a.StringArrMap[v_name]) <= label {
		fmt.Println("NameTableError:Not enougn length....")
		return "", false
	}
	return a.StringArrMap[v_name][label], true
}
func (a *NameTable) GetIntArr(v_name string, label int) (int, bool) {
	if len(a.IntArrMap[v_name]) <= label {
		fmt.Println("NameTableError:Not enougn length....")
		return 0, false
	}
	return a.IntArrMap[v_name][label], true
}
func (a *NameTable) GetBoolArr(v_name string, label int) (bool, bool) {
	if len(a.BoolArrMap[v_name]) <= label {
		fmt.Println("NameTableError:Not enougn length....")
		return false, false
	}
	return a.BoolArrMap[v_name][label], true
}

//向符号表中的对应名称变量和对应类型的变量进行赋值
func (a *NameTable) SetString(v_name string, v_value string) bool {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	}
	if v_type != "string" {
		fmt.Println("NameTableError:Type mismatch....")
		return false
	}
	a.StringMap[v_name] = v_value
	return true
}
func (a *NameTable) SetInt(v_name string, v_value int) bool {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	}
	if v_type != "int" {
		fmt.Println("NameTableError:Type mismatch....")
		return false
	}
	a.IntMap[v_name] = v_value
	return true
}
func (a *NameTable) SetBool(v_name string, v_value bool) bool {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	}
	if v_type != "bool" {
		fmt.Println("NameTableError:Type mismatch....")
		return false
	}
	a.BoolMap[v_name] = v_value
	return true
}

//向符号表中中对应变量数组进行赋值（分多个类型进行处理）
func (a *NameTable) SetStringArr(v_name string, label int, v_value string) bool {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	}
	if v_type != "stringarr" {
		fmt.Println("NameTableError:Type mismatch....")
		return false
	}
	if len(a.StringArrMap[v_name]) <= label {
		fmt.Println("NameTableError:Not enougn length....")
		return false
	}
	a.StringArrMap[v_name][label] = v_value
	return true
}
func (a *NameTable) SetIntArr(v_name string, label int, v_value int) bool {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	}
	if v_type != "intarr" {
		fmt.Println("NameTableError:Type mismatch....")
		return false
	}
	if len(a.IntArrMap[v_name]) <= label {
		fmt.Println("NameTableError:Not enougn length....")
		return false
	}
	a.IntArrMap[v_name][label] = v_value
	return true
}
func (a *NameTable) SetBoolArr(v_name string, label int, v_value bool) bool {
	v_type, ok := a.TypeMap[v_name]
	if ok == false {
		fmt.Println("NameTableError : There isn't a variable named " + v_name)
		return false
	}
	if v_type != "boolarr" {
		fmt.Println("NameTableError:Type mismatch....")
		return false
	}
	if len(a.BoolArrMap[v_name]) <= label {
		fmt.Println("NameTableError:Not enougn length....")
		return false
	}
	a.BoolArrMap[v_name][label] = v_value
	return true
}

func (a *NameTable) Show() {
	fmt.Println("========================SHOW NAMETABLE==============================")
	for k, v := range a.TypeMap {
		fmt.Print(k + " : " + v)
		switch v {
		case "int":
			fmt.Print("=")
			fmt.Println(a.IntMap[k])
		case "bool":
			fmt.Print("=")
			fmt.Println(a.BoolMap[k])
		case "string":
			fmt.Print("=")
			fmt.Println(a.StringMap[k])
		case "intarr":
			fmt.Println(len(a.IntArrMap[k]))
			fmt.Println(a.IntArrMap[k])
		case "boolarr":
			fmt.Println(len(a.BoolArrMap[k]))
			fmt.Println(a.BoolArrMap[k])
		case "stringarr":
			fmt.Println(len(a.StringArrMap[k]))
			fmt.Println(a.StringArrMap[k])
		default:
			fmt.Println("ERROR!!!!!!!!!!!!!!!!!!!!!")
		}
	}
	fmt.Println("========================SHOW NAMETABLE==============================")
}

