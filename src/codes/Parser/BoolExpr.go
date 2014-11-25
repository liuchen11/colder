package parser

/***************************************************************************************
**文件名：BoolExpr.go
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

//类：BoolExpr
type BoolExpr struct {
	boolexpr_expr *Expr  //表达式的内容
	boolexpr_type string //BoolExpr是否为空

	BoolExpr_value string //BoolExpr的值，为false 或者 true
}

//BoolExprError报错
func BoolExprError() {
	fmt.Println("There is an error in boolExpr")
}

//BoolExprError运行时报错
func BoolExprRunError() {
	fmt.Println("RunTimeError:There is an error in BoolExpr")
}

//生成一个指针BoolExpr
func GetBoolExpr(tokenlist []Token) (*BoolExpr, bool) {
	fmt.Print("BOOLEXPR:")
	showTokenlist(tokenlist)
	ret := new(BoolExpr)
	size := len(tokenlist)
	if size == 0 {
		//第一种类型：一个空的BoolExpr
		ret.boolexpr_type = "empty"
		return ret, true
	} else {
		//第二种类型:一个Expr
		ret.boolexpr_type = "expr"
		flag := false
		ret.boolexpr_expr, flag = GetExpr(tokenlist)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			BoolExprError()
			return nil, false
		}
		return ret, true
	}
	return ret, true
}

//BoolExpr执行
func (a *BoolExpr) Exe(na *NameTable, fu *FuncTable) bool {
	if a.boolexpr_type == "empty" {
		//如果为空，直接跳过
		a.BoolExpr_value = "true"
		return true
	} else {
		//如果不为空，执行Expr
		flag := a.boolexpr_expr.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			BoolExprRunError()
			return false
		}
		//结果必须为bool
		if a.boolexpr_expr.Expr_value_type != "bool" {
			BoolExprError()
			return false
		}
		//进行赋值
		a.BoolExpr_value = a.boolexpr_expr.Expr_value
		return true
	}
}

//调试使用：输出BoolExpr的结果
func (a *BoolExpr) Show() {
	fmt.Println(a.BoolExpr_value)
}
