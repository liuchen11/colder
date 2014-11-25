package tools

/***************************************************************************************
**文件名：Queue.go
**包名称：tools
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


type Queue struct {
	elements *List
	size     int
}

func NewQueue() *Queue {
	ret := new(Queue)
	ret.elements = NewList()
	ret.size = 0
	return ret
}

func (a *Queue) Push_Front(value any) {
	a.elements.Insert(0, value)
	a.size++
}

func (a *Queue) Push_Back(value any) {
	a.elements.Insert(a.size, value)
	a.size++
}

func (a *Queue) Pop_Front() bool {
	if a.size == 0 {
		return false
	} else {
		a.elements.DeleteByIndex(1)
		a.size--
		return true
	}
}

func (a *Queue) Pop_Back() bool {
	if a.size == 0 {
		return false
	} else {
		a.elements.DeleteByIndex(a.size)
		a.size--
		return true
	}
}

func (a *Queue) Head() any {
	return a.elements.Get(1)
}

func (a *Queue) Tail() any {
	return a.elements.Get(a.size)
}

func (a *Queue) Size() int {
	return a.size
}

func (a *Queue) Clear() {
	a.elements.Clear()
	a.size = 0
}

