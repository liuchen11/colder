package interpreter

/***************************************************************************************
**文件名：IfStmt.go
**包名称：interpreter
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"fmt"
)

type IfStmt struct {
}

func NewIfStmt() *IfStmt {
	ret := new(IfStmt)
	return ret
}

func (now IfStmt) Execute(a []Token, name *NameTable) ([]Token, bool) {
	fmt.Println("==========================BEGIN TOKEN==========================")
	size := len(a)
	i := 0
	b := make([]Token, 0, 5)
	if a[i].Content == "if" && a[i].Type == "Keyword" {
		i++
		if a[i].Content == "(" && a[i].Type == "Operator" {
			i++
			tmp := 0
			b = make([]Token, 0, 5)
			for {
				if i == size {
					fmt.Println("There is an error in IfStmt")
					return nil, false
				}
				if a[i].Content == ")" && a[i].Type == "Operator" && tmp == 0 {
					//fmt.Print("Largelymfs")
					//fmt.Println(i)
					break
				}
				if a[i].Content == "(" && a[i].Type == "Operator" {
					tmp++
				} else {
					if a[i].Content == ")" && a[i].Type == "Operator" {
						tmp--
					}
				}
				b = append(b, a[i])
				i++
			}
			fmt.Println(b)
			expr := NewExpr()
			v_type, v_value := expr.GetValue(b, name)
			if v_type != "bool" {
				fmt.Println("There is an error in IfStmt")
				return nil, false
			}
			if v_value == "true" {
				//fmt.Println("CHECK")
				i++
				if a[i].Content == "{" && a[i].Type == "Operator" {
					tmp = 0
					b = make([]Token, 0, 5)
					i++
					for {
						if i == size {
							fmt.Println("There is an error in IfStmt")
							return nil, false
						}
						if a[i].Content == "}" && a[i].Type == "Operator" && tmp == 0 {
							//fmt.Print("Largelymfs")
							//fmt.Println(i)
							break
						}
						if a[i].Content == "{" && a[i].Type == "Operator" {
							tmp++
						} else {
							if a[i].Content == "}" && a[i].Type == "Operator" {
								tmp--
							}
						}
						b = append(b, a[i])
						i++
					}
					i++
					fmt.Println(b)
					//BlockStmt
					//BlockStmt(b, name)
					//Skip Else
					//fmt.Println(a[i])
					if a[i].Content == "else" && a[i].Type == "Keyword" {

						i++
						tmp = 0
						//fmt.Println(a[i].Content)
						if a[i].Content == "{" && a[i].Type == "Operator" {
							i++
							for {
								if i == size {
									fmt.Println("There is an error in IfStmt")
									return nil, false
								}
								if a[i].Content == "}" && a[i].Type == "Operator" && tmp == 0 {
									break
								}
								if a[i].Content == "{" && a[i].Type == "Operator" {
									tmp++
								} else {
									if a[i].Content == "}" && a[i].Type == "Operator" {
										tmp--
									}
								}
								i++
							}
							i++
							if i < size {
								b = a[i:]
							} else {
								b = nil
							}
						} else {
							fmt.Println("There is an error in IfStmt")
							return nil, false
						}

					} else {
						b = a[i:]
					}
				} else {
					fmt.Println("There is an error in IfStmt")
					return nil, false
				}
			} else {
				//SkipBlockStmt
				//Else
			}
		} else {
			fmt.Println("There is an error in IfStmt")
			return nil, false
		}
	} else {
		fmt.Println("There is an error in IfStmt...")
		return nil, false
	}
	fmt.Println("===========================END TOKEN===========================")
	fmt.Println(b)
	return b, true
}

