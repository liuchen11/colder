package interpreter

/***************************************************************************************
**文件名：Token.go
**包名称：interpreter
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


type Token struct {
	Content  string //内容
	Type     string //类型 Keyword Reservedword Identifier Int Operator
	Index    int    //关键字保留字编号
	IntValue int    //整型字值
	X        int    //列号
	Y        int    //行号
}

