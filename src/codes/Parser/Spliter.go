package parser

/***************************************************************************************
**文件名：Spliter.go
**包名称：parser
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"bufio"
	"codes/tools"
	"fmt"
	"io"
	"os"
	"strings"
)

func SplitFile(filename string) bool {
	fin, err := os.Open(filename)
	if err != nil {
		fmt.Println("there are some problems opening the file")
		return false
	}
	defer fin.Close()
	reader := bufio.NewReader(fin)
	buffer, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("there are some problems loading the file")
		return false
	}
	buf := string(buffer)
	buf = tools.WipeOutNote(buf)
	buf = tools.WipeOutBlankPrefix(buf)
	for {
		switch {
		case strings.Contains(buf, "<config>"):
			//fmt.Println("IN")
			start := strings.Index(buf, "<config>") + 8
			buf = buf[start:]
			config_writer, configerr := os.Create(filename + ".config")
			if configerr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer config_writer.Close()
			for {
				//fmt.Println("BEGIN")
				if strings.Contains(buf, "</config>") {
					end := strings.Index(buf, "</config>")
					config_writer.WriteString(buf[:end])
					end = end + 9
					buf = buf[end:]
					break
				} else {
					//fmt.Println("OKOK")
					if tools.IsBlank(buf) == false {
						config_writer.WriteString(buf + "\n")
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		case strings.Contains(buf, "<var>"):
			start := strings.Index(buf, "<var>") + 5
			buf = buf[start:]
			var_writer, varerr := os.Create(filename + ".var")
			if varerr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer var_writer.Close()
			for {
				if strings.Contains(buf, "</var>") {
					end := strings.Index(buf, "</var>")
					var_writer.WriteString(buf[:end])
					end = end + 6
					buf = buf[end:]
					break
				} else {
					if tools.IsBlank(buf) == false {
						var_writer.WriteString(buf + "\n")
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		case strings.Contains(buf, "<card>"):
			start := strings.Index(buf, "<card>") + 6
			buf = buf[start:]
			card_writer, carderr := os.Create(filename + ".card")
			if carderr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer card_writer.Close()
			for {
				if strings.Contains(buf, "</card>") {
					end := strings.Index(buf, "</card>")
					card_writer.WriteString(buf[:end])
					end = end + 7
					buf = buf[end:]
					break
				} else {
					if tools.IsBlank(buf) == false {
						card_writer.WriteString(buf + "\n")
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		case strings.Contains(buf, "<mode>"):
			start := strings.Index(buf, "<mode>") + 6
			buf = buf[start:]
			mode_writer, modeerr := os.Create(filename + ".mode")
			if modeerr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer mode_writer.Close()
			var blank bool = true
			for {
				if strings.Contains(buf, "</mode>") {
					end := strings.Index(buf, "</mode>")
					mode_writer.WriteString(buf[:end])
					end = end + 7
					buf = buf[end:]
					break
				} else {
					buf = tools.WipeOutBlankString(buf)
					if tools.IsBlank(buf) {
						if blank == false {
							mode_writer.WriteString("\n")
							blank = true
						}
					} else {
						mode_writer.WriteString(buf)
						blank = false
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		case strings.Contains(buf, "<func>"):
			start := strings.Index(buf, "<func>") + 6
			buf = buf[start:]
			func_writer, funcerr := os.Create(filename + ".func")
			if funcerr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer func_writer.Close()
			for {
				if strings.Contains(buf, "</func>") {
					end := strings.Index(buf, "</func>")
					func_writer.WriteString(buf[:end])
					end = end + 7
					buf = buf[end:]
					break
				} else {
					if tools.IsBlank(buf) == false {
						func_writer.WriteString(buf + "\n")
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		case strings.Contains(buf, "<body>"):
			start := strings.Index(buf, "<body>") + 6
			buf = buf[start:]
			body_writer, bodyerr := os.Create(filename + ".body")
			if bodyerr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer body_writer.Close()
			for {
				if strings.Contains(buf, "</body>") {
					end := strings.Index(buf, "</body>")
					body_writer.WriteString(buf[:end])
					end = end + 7
					buf = buf[end:]
					break
				} else {
					if tools.IsBlank(buf) == false {
						body_writer.WriteString(buf + "\n")
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		case strings.Contains(buf, "<end>"):
			start := strings.Index(buf, "<end>") + 5
			buf = buf[start:]
			end_writer, enderr := os.Create(filename + ".end")
			if enderr != nil {
				fmt.Println("there are some problems creating the file")
				return false
			}
			defer end_writer.Close()
			for {
				if strings.Contains(buf, "</end>") {
					end := strings.Index(buf, "</end>")
					end_writer.WriteString(buf[:end])
					end = end + 6
					buf = buf[end:]
					break
				} else {
					if tools.IsBlank(buf) == false {
						end_writer.WriteString(buf + "\n")
					}
					buffer, _, err = reader.ReadLine()
					switch err {
					case io.EOF:
						return true
					case nil:
						buf = string(buffer)
						buf = tools.WipeOutNote(buf)
						buf = tools.WipeOutBlankPrefix(buf)
					default:
						fmt.Println("there are some problems loading the file")
						return false
					}
				}
			}
		default:
			buffer, _, err = reader.ReadLine()
			switch err {
			case io.EOF:
				return true
			case nil:
				buf = string(buffer)
				buf = tools.WipeOutNote(buf)
				buf = tools.WipeOutBlankPrefix(buf)
			default:
				fmt.Println("there are some problems loading the file")
				return false
			}
		}
	}
}

