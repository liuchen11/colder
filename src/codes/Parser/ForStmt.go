package parser

/***************************************************************************************
**文件名：ForStmt.go
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

//类ForStmt
type ForStmt struct {
	forstmt_simp1     *SimpleStmt
	forstmt_condition *BoolExpr
	forstmt_simp2     *SimpleStmt
	forstmt_content   *Stmt
}

//报错
func ForStmtError() {
	fmt.Println("There is an error in ForStmt....")
}

//生成函数
func GetForStmt(tokenlist []Token) (*ForStmt, bool) {
	fmt.Print("FORSTMT:")
	showTokenlist(tokenlist)
	ret := new(ForStmt)
	size := len(tokenlist)
	//不足最小的Token数
	if size < 5 {
		ForStmtError()
		return nil, false
	}
	//进行格式判断
	if tokenlist[0].Content == "for" && tokenlist[0].Type == "Keyword" {
		if tokenlist[1].Content == "(" && tokenlist[1].Type == "Operator" {
			i := 2
			//找到第一个;
			for i = 2; i < size; i++ {
				if tokenlist[i].Content == ";" && tokenlist[i].Type == "Operator" {
					break
				}
			}
			//如果没有找到
			if i == size {
				ForStmtError()
				return nil, false
			}
			//解析第一个SimpeStmt
			tmp := tokenlist[2:i]
			flag := false
			ret.forstmt_simp1, flag = GetSimpleStmt(tmp)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ForStmtError()
				return nil, false
			}
			//解析第二个BoolExpr
			j := 1
			//寻找下一个；
			for j = i + 1; j < size; j++ {
				if tokenlist[j].Content == ";" && tokenlist[j].Type == "Operator" {
					break
				}
			}
			//如果没有找到
			if j == size {
				ForStmtError()
				return nil, false
			}
			//截取下一个SimpleStmt
			tmp = tokenlist[i+1 : j]
			ret.forstmt_condition, flag = GetBoolExpr(tmp)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ForStmtError()
				return nil, false
			}
			i = j
			j = i + 1
			//寻找下一个SimpleStmt，结束符是）
			for j = i + 1; j < size; j++ {
				if tokenlist[j].Content == ")" && tokenlist[j].Type == "Operator" {
					break
				}
			}
			//如果没有找到
			if j == size {
				ForStmtError()
				return nil, false
			}
			tmp = tokenlist[i+1 : j]
			ret.forstmt_simp2, flag = GetSimpleStmt(tmp)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ForStmtError()
				return nil, false
			}
			j++
			tmp = tokenlist[j:size]
			//后面一个为一个语句
			//showTokenlist(tmp)
			ret.forstmt_content, flag = GetStmt(tmp)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ForStmtError()
				return nil, false
			}
			return ret, true
		} else {
			ForStmtError()
			return nil, false
		}
	} else {
		ForStmtError()
		return nil, false
	}
	return ret, true
}

//报错
func ForStmtRunError() {
	fmt.Println("RunTimeError : There is an error in Forstmt")
}

//运行
func (a *ForStmt) Exe(na *NameTable, fu *FuncTable) bool {
	//开始时候执行SimpleStmt1
	flag := a.forstmt_simp1.Exe(na, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		ForStmtRunError()
		return false
	}
	for {
		//1.执行BoolExpr
		//2.根据BoolExpr的真假进行判断
		//3.如果为真进行循环，否则结束
		//4.执行Content
		//5.执行第二个SimpleStmt
		//5.跳回1
		flag = a.forstmt_condition.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ForStmtRunError()
			return false
		}
		if a.forstmt_condition.BoolExpr_value == "false" {
			//跳出循环
			break
		}
		//执行Content
		flag = a.forstmt_content.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ForStmtRunError()
			return false
		}
		//执行后一句SimpleStmt
		flag = a.forstmt_simp2.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ForStmtRunError()
			return false
		}
	}
	return true
}
