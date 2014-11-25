package interpreter

/***************************************************************************************
**文件名：Expr.go
**包名称：interpreter
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"fmt"
	"strconv"
)

type Expr struct {
}

func NewExpr() *Expr {
	ret := new(Expr)
	return ret
}
func (r *Expr) GetValue(a []Token, name *NameTable) (string, string) {
	v_type, v_value := r.GetExpr(a, name)
	return v_type, v_value
}
func (r *Expr) GetExpr(a []Token, name *NameTable) (string, string) {
	tmp := 0
	size := len(a)
	i := 0
	for i = 0; i < size; i++ {
		if tmp == 0 && a[i].Type == "Operator" && (a[i].Content == "&&" || a[i].Content == "||" || a[i].Content == "==" || a[i].Content == "!=" || a[i].Content == ">" || a[i].Content == "<" || a[i].Content == ">=" || a[i].Content == "<=") {
			fmt.Println("Largelymfs")
			break
		} else {
			if a[i].Type == "Operator" && a[i].Content == "(" {
				tmp = tmp + 1
			} else {
				if a[i].Type == "Operator" && a[i].Content == ")" {
					tmp = tmp - 1
				}
			}
		}
	}
	fmt.Println(i)
	if i == size {
		return r.GetOpt(a, name)
	} else {
		left := a[0:i]
		right := a[i+1:]
		left_type, left_value := r.GetOpt(left, name)
		right_type, right_value := r.GetExpr(right, name)
		if left_type == "error" || right_type == "error" {
			return "error", ""
		} else {
			if left_type != right_type {
				if left_type == "int" && right_type == "float" {
					if (a[i].Content == "&&" && a[i].Type == "Operator") || (a[i].Type == "Operator" && a[i].Content == "||") {
						return "error", ""
					}
					v_type := "bool"
					v_value := ""
					left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
					left_value_float := float64(left_value_int)
					right_value_float, _ := strconv.ParseFloat(right_value, 64)
					switch a[i].Content {
					case "==":
						if left_value_float == right_value_float {
							v_value = "true"
						} else {
							v_value = "false"
						}
					case "!=":
						if left_value_float != right_value_float {
							v_value = "true"
						} else {
							v_value = "false"
						}
					case ">=":
						if left_value_float >= right_value_float {
							v_value = "true"
						} else {
							v_value = "false"
						}
					case "<=":
						if left_value_float <= right_value_float {
							v_value = "true"
						} else {
							v_value = "false"
						}
					case ">":
						if left_value_float > right_value_float {
							v_value = "true"
						} else {
							v_value = "false"
						}
					case "<":
						if left_value_float < right_value_float {
							v_value = "true"
						} else {
							v_value = "false"
						}
					}
					return v_type, v_value
				} else {
					if left_type == "float" && right_type == "int" {
						if (a[i].Content == "&&" && a[i].Type == "Operator") || (a[i].Type == "Operator" && a[i].Content == "||") {
							return "error", ""
						}
						v_type := "bool"
						v_value := ""
						right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
						right_value_float := float64(right_value_int)
						left_value_float, _ := strconv.ParseFloat(left_value, 64)
						switch a[i].Content {
						case "==":
							if left_value_float == right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "!=":
							if left_value_float != right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case ">=":
							if left_value_float >= right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "<=":
							if left_value_float <= right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case ">":
							if left_value_float > right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "<":
							if left_value_float < right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						}
						return v_type, v_value
					} else {
						return "error", ""
					}
				}
			} else {
				v_value := ""
				v_type := ""
				switch a[i].Content {
				case "==", "!=":
					switch a[i].Content {
					case "==":
						switch left_type {
						case "string":
							v_type = "bool"
							if left_value == right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "int":
							v_type = "bool"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							if left_value_int == right_value_int {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "float":
							v_type = "bool"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							if left_value_float == right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "bool":
							v_type = "bool"
							if left_value == right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						default:
							v_value = ""
							v_type = "error"
						}
					case "!=":
						switch left_type {
						case "string":
							v_type = "bool"
							if left_value != right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "int":
							v_type = "bool"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							if left_value_int != right_value_int {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "float":
							v_type = "bool"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							if left_value_float != right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "bool":
							v_type = "bool"
							if left_value != right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						default:
							v_value = ""
							v_type = "error"
						}
					}
				case ">=", "<=", ">", "<":
					switch a[i].Content {
					case ">=":
						switch left_type {
						case "string":
							v_type = "bool"
							if left_value >= right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "int":
							v_type = "bool"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							if left_value_int >= right_value_int {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "float":
							v_type = "bool"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							if left_value_float >= right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						default:
							v_value = ""
							v_type = "error"
						}
					case "<=":
						switch left_type {
						case "string":
							v_type = "bool"
							if left_value <= right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "int":
							v_type = "bool"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							if left_value_int <= right_value_int {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "float":
							v_type = "bool"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							if left_value_float <= right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						default:
							v_value = ""
							v_type = "error"
						}
					case ">":
						switch left_type {
						case "string":
							v_type = "bool"
							if left_value > right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "int":
							v_type = "bool"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							if left_value_int > right_value_int {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "float":
							v_type = "bool"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							if left_value_float > right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						default:
							v_value = ""
							v_type = "error"
						}
					case "<":
						switch left_type {
						case "string":
							v_type = "bool"
							if left_value < right_value {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "int":
							v_type = "bool"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							if left_value_int < right_value_int {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "float":
							v_type = "bool"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							if left_value_float < right_value_float {
								v_value = "true"
							} else {
								v_value = "false"
							}
						default:
							v_value = ""
							v_type = "error"
						}
					}
				case "&&", "||":
					if left_type != "bool" {
						return "error", ""
					} else {
						v_type = "bool"
						switch a[i].Content {
						case "&&":
							if left_value == "true" && right_value == "true" {
								v_value = "true"
							} else {
								v_value = "false"
							}
						case "||":
							if left_value == "true" || right_value == "true" {
								v_value = "true"
							} else {
								v_value = "false"
							}
						}
					}
				}
				return v_type, v_value
			}
		}
	}
	return "error", ""
}
func (r *Expr) GetOpt(a []Token, name *NameTable) (string, string) {
	size := len(a)
	i := 0
	tmp := 0
	for i = size - 1; i >= 0; i-- {
		if tmp == 0 && a[i].Type == "Operator" && (a[i].Content == "+" || a[i].Content == "-") {
			break
		} else {
			if a[i].Type == "Operator" && a[i].Content == "(" {
				tmp = tmp - 1
			} else {
				if a[i].Type == "Operator" && a[i].Content == ")" {
					tmp = tmp + 1
				}
			}
		}
	}
	if i == -1 {
		return r.GetTerm(a, name)
	} else {
		v_value := ""
		v_type := ""
		left := a[0:i]
		right := a[i+1:]
		left_type, left_value := r.GetOpt(left, name)
		right_type, right_value := r.GetTerm(right, name)
		if left_type == "error" || right_type == "error" {
			return "error", ""
		} else {
			if left_type != right_type {
				if left_type == "int" && right_type == "float" {
					switch a[i].Content {
					case "+":
						v_type = "float"
						left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
						left_value_float := float64(left_value_int)
						right_value_float, _ := strconv.ParseFloat(right_value, 64)
						v_value = strconv.FormatFloat(left_value_float+right_value_float, 'f', 3, 64)
					case "-":
						v_type = "float"
						left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
						left_value_float := float64(left_value_int)
						right_value_float, _ := strconv.ParseFloat(right_value, 64)
						v_value = strconv.FormatFloat(left_value_float-right_value_float, 'f', 3, 64)
					}
					return v_type, v_value
				} else {
					if left_type == "float" && right_type == "int" {
						switch a[i].Content {
						case "+":
							v_type = "float"
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							right_value_float := float64(right_value_int)
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							v_value = strconv.FormatFloat(left_value_float+right_value_float, 'f', 3, 64)
						case "-":
							v_type = "float"
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							right_value_float := float64(right_value_int)
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							v_value = strconv.FormatFloat(left_value_float-right_value_float, 'f', 3, 64)
						}
						return v_type, v_value
					} else {
						return "error", ""
					}
				}
			} else {
				if left_type == "string" || left_type == "int" || left_type == "float" {
					switch a[i].Content {
					case "+":
						switch left_type {
						case "string":
							v_type = "string"
							v_value = left_value + right_value
						case "int":
							v_type = "int"
							//v_value = strconv.FormatInt(int64(strconv.ParseInt(left_value,10, 32)+strconv.ParseInt(right_value,10,32)),10)
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							v_value = strconv.FormatInt(int64(left_value_int+right_value_int), 10)
						case "float":
							v_type = "float"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							v_value = strconv.FormatFloat(left_value_float+right_value_float, 'f', 3, 64)
						}
					case "-":
						switch left_type {
						case "string":
							v_type = "error"
							v_value = ""
						case "int":
							v_type = "int"
							left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							v_value = strconv.FormatInt(int64(left_value_int-right_value_int), 10)
						case "float":
							v_type = "float"
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							right_value_float, _ := strconv.ParseFloat(right_value, 64)
							v_value = strconv.FormatFloat(left_value_float-right_value_float, 'f', 3, 64)
						}
					}
					return v_type, v_value
				} else {
					return "error", ""
				}
			}
		}
	}
	return "error", ""
}
func (r *Expr) GetTerm(a []Token, name *NameTable) (string, string) {
	size := len(a)
	i := 0
	tmp := 0
	for i = size - 1; i >= 0; i-- {
		if tmp == 0 && a[i].Type == "Operator" && (a[i].Content == "*" || a[i].Content == "/") {
			break
		} else {
			if a[i].Type == "Operator" && a[i].Content == "(" {
				tmp = tmp - 1
			} else {
				if a[i].Type == "Operator" && a[i].Content == ")" {
					tmp = tmp + 1
				}
			}
		}
	}
	fmt.Print("Largelymfs___")
	fmt.Println(i)
	if i == -1 {
		return r.GetFactor(a, name)
	} else {
		v_value := ""
		v_type := ""
		left := a[0:i]
		right := a[i+1:]
		fmt.Println(left)
		fmt.Println(right)
		left_type, left_value := r.GetTerm(left, name)
		right_type, right_value := r.GetFactor(right, name)
		if left_type == "error" || right_type == "error" {
			return "error", ""
		} else {
			if left_type != right_type {
				if left_type == "int" && right_type == "float" {
					switch a[i].Content {
					case "*":
						v_type = "float"
						left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
						left_value_float := float64(left_value_int)
						right_value_float, _ := strconv.ParseFloat(right_value, 64)
						v_value = strconv.FormatFloat(left_value_float*right_value_float, 'f', 3, 64)
					case "/":
						v_type = "float"
						left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
						left_value_float := float64(left_value_int)
						right_value_float, _ := strconv.ParseFloat(right_value, 64)
						if right_value_float < 0.001 {
							return "error", ""
						}
						v_value = strconv.FormatFloat(left_value_float/right_value_float, 'f', 3, 64)
					}
					return v_type, v_value
				} else {
					if left_type == "float" && right_type == "int" {
						switch a[i].Content {
						case "*":
							v_type = "float"
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							right_value_float := float64(right_value_int)
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							v_value = strconv.FormatFloat(left_value_float*right_value_float, 'f', 3, 64)
						case "/":
							v_type = "float"
							right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
							right_value_float := float64(right_value_int)
							left_value_float, _ := strconv.ParseFloat(left_value, 64)
							if right_value_float < 0.001 {
								return "error", ""
							}
							v_value = strconv.FormatFloat(left_value_float/right_value_float, 'f', 3, 64)
						}
						return v_type, v_value
					} else {
						return "error", ""
					}
				}
			} else {
				switch a[i].Content {
				case "*":
					switch left_type {
					case "string":
						v_value = ""
						v_type = "error"
					case "int":
						v_type = "int"
						left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
						right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
						v_value = strconv.FormatInt(int64(left_value_int*right_value_int), 10)
					case "float":
						v_type = "float"
						left_value_float, _ := strconv.ParseFloat(left_value, 64)
						right_value_float, _ := strconv.ParseFloat(right_value, 64)
						v_value = strconv.FormatFloat(left_value_float*right_value_float, 'f', 3, 64)
					case "bool":
						v_type = "error"
						v_value = ""
					case "error":
						v_type = "error"
						v_value = ""
					}
				case "/":
					switch left_type {
					case "string":
						v_value = ""
						v_type = "error"
					case "int":
						v_type = "int"
						left_value_int, _ := strconv.ParseInt(left_value, 10, 32)
						right_value_int, _ := strconv.ParseInt(right_value, 10, 32)
						if right_value_int == 0 {
							return "error", ""
						}
						v_value = strconv.FormatInt(int64(left_value_int/right_value_int), 10)
					case "float":
						v_type = "float"
						left_value_float, _ := strconv.ParseFloat(left_value, 64)
						right_value_float, _ := strconv.ParseFloat(right_value, 64)
						if right_value_float < 0.001 {
							return "error", ""
						}
						v_value = strconv.FormatFloat(left_value_float/right_value_float, 'f', 3, 64)
					case "bool":
						v_type = "error"
						v_value = ""
					case "error":
						v_type = "error"
						v_value = ""
					}
				}
				return v_type, v_value
			}
		}
	}
	return "error", ""
}
func (r *Expr) GetFactor(a []Token, name *NameTable) (string, string) {
	size := len(a)
	fmt.Println(size)
	if size == 1 {
		switch a[0].Type {
		case "Int":
			return "int", a[0].Content
		case "Float":
			return "float", a[0].Content
		case "Keyword":
			if a[0].Content == "false" || a[0].Content == "true" {
				return "bool", a[0].Content
			} else {
				return "error", ""
			}
		case "Identifier":
			now_type := name.GetType(a[0].Content)
			switch now_type {
			case "string":
				return "string", name.GetString(a[0].Content)
			case "bool":
				tmp := name.GetBool(a[0].Content)
				if tmp == true {
					return "bool", "true"
				} else {
					return "bool", "false"
				}
			case "int":
				return "int", strconv.FormatInt(int64(name.GetInt(a[0].Content)), 10)
			case "float":
				return "float", strconv.FormatFloat(name.GetFloat(a[0].Content), 'f', 3, 64)
			case "error":
				return "error", ""
			}
		default:
			return "error", ""
		}
	} else {
		if a[0].Content == "(" && a[0].Type == "Operator" && a[size-1].Content == ")" && a[size-1].Type == "Operator" {
			b := a[1 : size-1]
			return r.GetExpr(b, name)
		} else {
			if size == 3 {
				if a[0].Type == "Operator" && a[0].Content == "\"" && a[2].Content == "\"" && a[2].Type == "Operator" {
					return "string", a[1].Content
				} else {
					return "error", ""
				}
			} else {
				return "error", ""
			}
		}
	}
	return "error", ""
}

