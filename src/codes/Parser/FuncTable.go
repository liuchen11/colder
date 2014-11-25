package parser

/***************************************************************************************
**文件名：FuncTable.go
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

//函数调用表
type FuncTable struct {
	func_content []*FuncDef
	TypeMap      map[string]string
	ValueMap     map[string]string

	now_func []string
	cur      int
}

//报错
func FuncTableError() {
	fmt.Println("There is an error in FuncTable...")
}

//生成函数
func GetFuncTable(tokenlist []Token) (*FuncTable, bool) {
	fmt.Print("FUNCTABLE:")
	showTokenlist(tokenlist)
	ret := new(FuncTable)
	ret.TypeMap = make(map[string]string)
	ret.ValueMap = make(map[string]string)
	ret.now_func = make([]string, 0, 5)
	ret.cur = 0
	size := len(tokenlist)
	ret.func_content = make([]*FuncDef, 0, 5)
	if size == 0 {
		return nil, true
	}
	i := 0
	tmp1 := 0
	j := i
	for {
		for j = i; j < size; j++ {
			if tmp1 == 1 && tokenlist[j].Content == "}" && tokenlist[j].Type == "Operator" {
				break
			}
			if tmp1 == 0 && tokenlist[j].Content == ";" && tokenlist[j].Type == "Operator" {
				break
			}
			if tokenlist[j].Content == "{" && tokenlist[j].Type == "Operator" {
				tmp1++
			}
			if tokenlist[j].Content == "}" && tokenlist[j].Type == "Operator" {
				tmp1--
			}
		}
		tmp := tokenlist[i : j+1]
		next_func, flag := GetFuncDef(tmp)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			FuncTableError()
			return nil, false
		}
		i = j + 1
		tmp1 = 0
		ret.func_content = append(ret.func_content, next_func)
		fmt.Println("FUNCTION :" + next_func.Funcdef_name + " : " + next_func.Funcdef_type)
		if j == size-1 {
			break
		}
	}
	return ret, true
}

//添加一个函数
func (a *FuncTable) AddFunc(v_name string) {
	size := len(a.now_func)
	if size == a.cur {
		a.now_func = append(a.now_func, v_name)
		a.cur++
		return
	}
	a.now_func[a.cur] = v_name
	a.cur++
}

//处理一个reuturn语句
func (a *FuncTable) ReturnFunc(v_type string, v_value string) bool {
	//now_type := a.TypeMap[a.now_func[a.cur-1]]
	//if now_type != v_type {
	//	fmt.Println("Type Mismatch")
	//	return false
	//}
	a.ValueMap[a.now_func[a.cur-1]] = v_value
	a.cur--
	return true
}

//获得调用栈的最上面的type和value
func (a *FuncTable) Get(v_name string) (string, string) {
	return a.TypeMap[v_name], a.ValueMap[v_name]
}

//执行
func (a *FuncTable) Exe(name string, na *NameTable, in_type []string, in_value []string) bool {
	size := len(a.func_content)
	for i := 0; i < size; i++ {
		//找到名称相同的一项
		if name == a.func_content[i].Funcdef_name {
			flag := a.func_content[i].Exe(na, a, in_type, in_value)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				FuncTableError()
				return false
			}
			//每次执行一个Func，更新到TypeMap中
			a.TypeMap[name] = a.func_content[i].Funcdef_type
		}
	}
	return true
}
