package tools

/***************************************************************************************
**文件名：Stack.go
**包名称：tools
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


type Stack struct {
	size int
	item []any
}

func (a *Stack) Push(toadd any) {
	a.item = append(a.item[0:a.size], toadd)
	a.size++
}

func (a *Stack) Pop() bool {
	if a.size > 0 {
		a.size--
		return true
	} else {
		return false
	}
}

func (a *Stack) Top() any {
	if a.size > 0 {
		return a.item[a.size-1]
	} else {
		return nil
	}
}

func (a *Stack) Empty() bool {
	if a.size == 0 {
		return true
	} else {
		return false
	}
}

func (a *Stack) Clear() {
	a.size = 0
}

func (a *Stack) Size() int {
	return a.size
}

