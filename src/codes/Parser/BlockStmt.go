package parser

/***************************************************************************************
**文件名：BlockStmt.go
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

//类Block:语句块
type BlockStmt struct {
	blockstmt_content []*Stmt //语句块
}

//语句块报错——运行前
func BlockStmtError() {
	fmt.Println("There is an error in BlockStmt")
}

//语句块报错--运行中
func BlockStmtRunError() {
	fmt.Println("RunTimeError : There is an error in BlockStmt")
}

//获得一个BlockStmt的指针
func GetBlockStmt(tokenlist []Token) (*BlockStmt, bool) {
	fmt.Print("BLOCKSTMT:")
	showTokenlist(tokenlist)
	ret := new(BlockStmt)
	size := len(tokenlist)
	if size < 2 {
		BlockStmtError()
		return nil, false
	}

	//BlockStmt的两边应该是{}
	if tokenlist[0].Content == "{" && tokenlist[0].Type == "Operator" {
		if tokenlist[size-1].Content == "}" && tokenlist[size-1].Type == "Operator" {
			i := 1
			j := i
			tmp := 0
			tmp1 := 0
			label := 0
			for {
				tmp = 0
				if i >= size {
					BlockStmtError()
					return nil, false
				}
				j = i

				//寻找下一条指令
				for j = i; j < size; j++ {

					//终止结果，找到一个最外层的右括号
					if tmp1 == 0 && tmp == 0 && tokenlist[j].Content == "}" && tokenlist[j].Type == "Operator" {
						label = 1
						break
					}
					//遇到一个右括号，找到下一句
					if tmp1 == 0 && tmp == 1 && tokenlist[j].Content == "}" && tokenlist[j].Type == "Operator" {
						label = 2
						break
					}
					//遇到一个分好，找到下一句
					if tmp1 == 0 && tmp == 0 && tokenlist[j].Content == ";" && tokenlist[j].Type == "Operator" {
						label = 3
						break
					}
					//括号匹配
					if tokenlist[j].Type == "Operator" {
						switch tokenlist[j].Content {
						case "{":
							tmp++
						case "}":
							tmp--
						case "(":
							tmp1--
						case ")":
							tmp1++
						}
					}
				}
				if j == size {
					BlockStmtError()
					return nil, false
				}
				//遇到终止条件，退出
				if label == 1 {
					break
				}
				//新添加一个语句
				switch label {
				case 1:
				case 2:
					tmp := tokenlist[i : j+1]
					next_stmt, flag := GetStmt(tmp)
					//判断是否在下面出错，完成栈式报错
					if flag == false {
						BlockStmtError()
						return nil, false
					}
					ret.blockstmt_content = append(ret.blockstmt_content, next_stmt)
					i = j + 1
				case 3:
					tmp := tokenlist[i : j+1]
					next_stmt, flag := GetStmt(tmp)
					//判断是否在下面出错，完成栈式报错
					if flag == false {
						BlockStmtError()
						return nil, false
					}
					ret.blockstmt_content = append(ret.blockstmt_content, next_stmt)
					i = j + 1
				default:
					BlockStmtError()
					return nil, false
				}
			}
		} else {
			BlockStmtError()
			return nil, false
		}
	} else {
		BlockStmtError()
		return nil, false
	}
	return ret, true
}

//Block语句块执行
func (a *BlockStmt) Exe(na *NameTable, fu *FuncTable) bool {
	flag := false
	size := len(a.blockstmt_content)
	//每一条指令分别执行
	for i := 0; i < size; i++ {
		flag = a.blockstmt_content[i].Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			BlockStmtRunError()
			return false
		}
	}
	return true
}
