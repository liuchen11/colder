package tools

/***************************************************************************************
**文件名：List.go
**包名称：tools
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


type any interface{}

type node struct {
	pre   *node
	next  *node
	value any
}

func NewNode(v any, p *node, n *node) *node {
	ret := new(node)
	ret.value = v
	ret.pre = p
	ret.next = n
	return ret
}

type List struct {
	first *node
	last  *node
	size  int
}

func NewList() *List {
	ret := new(List)
	ret.size = 0
	ret.first = NewNode(nil, nil, ret.last)
	ret.last = NewNode(nil, ret.first, nil)
	return ret
}

//插入到第index个元素之后
//first的下标记为0
//返回值为插入操作是否合法
func (a *List) Insert(index int, value any) bool {
	if index > a.size {
		return false
	} else {
		present := a.first
		for i := 0; i < index; i++ {
			present = present.next
		}
		toinsert := NewNode(value, present, present.next)
		present.next.pre = toinsert
		present.next = toinsert
		a.size++
		return true
	}
}

//获取第index个元素
func (a *List) Get(index int) any {
	if index < 0 || index > a.size {
		return nil
	} else {
		present := a.first
		for i := 0; i < index; i++ {
			present = present.next
		}
		return present.value
	}
}

//按照下标删除其中的节点，成功返回true失败，返回false
func (a *List) DeleteByIndex(index int) bool {
	if index <= 0 || index > a.size {
		return false
	} else {
		present := a.first
		for i := 0; i < index; i++ {
			present = present.next
		}
		prenode := present.pre
		nextnode := present.next
		prenode.next = nextnode
		nextnode.pre = prenode
		a.size--
		return true
	}
}

//按照值删除其中的节点，同上
func (a *List) DeleteByValue(v any) bool {
	present := a.first.next
	for i := 1; i <= a.size; i++ {
		if present.value == v {
			prenode := present.pre
			nextnode := present.next
			prenode.next = nextnode
			nextnode.pre = prenode
			a.size--
			return true
		}
		present = present.next
	}
	return false
}

//寻找一个数
func (a *List) Find(v any) int {
	present := a.first.next
	for i := 1; i <= a.size; i++ {
		if present.value == v {
			return i
		}
		present = present.next
	}
	return -1
}

//元素多少
func (a *List) Size() int {
	return a.size
}

//完全清空
func (a *List) Clear() {
	a.first.next = a.last
	a.last.pre = a.first
	a.size = 0
}

