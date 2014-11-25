package interpreter

/***************************************************************************************
**文件名：Expr_test.go
**包名称：interpreter
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
	"os"
	"strconv"
	"testing"
)

func gernerate_Factor(file string, output *os.File) {
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := NewParser(br)
	p.SplitToken()
	p.Show()
	expr := NewExpr()
	name := NewNameTable()
	v_type, v_value := expr.GetFactor(p.Content(), name)
	//fmt.Println("TYPE : " + v_type)
	//fmt.Println("VALUE : " + v_value)
	output.WriteString("TYPE:" + v_type + " " + "VALUE:" + v_value + "\n")
}
func gernerate_Expr(file string, output *os.File) {
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := NewParser(br)
	p.SplitToken()
	p.Show()
	expr := NewExpr()
	name := NewNameTable()
	v_type, v_value := expr.GetExpr(p.Content(), name)
	//fmt.Println("TYPE : " + v_type)
	//fmt.Println("VALUE : " + v_value)
	output.WriteString("TYPE:" + v_type + " " + "VALUE:" + v_value + "\n")
}
func gernerate_Term(file string, output *os.File) {
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := NewParser(br)
	p.SplitToken()
	p.Show()
	expr := NewExpr()
	name := NewNameTable()
	v_type, v_value := expr.GetTerm(p.Content(), name)
	//fmt.Println("TYPE : " + v_type)
	//fmt.Println("VALUE : " + v_value)
	output.WriteString("TYPE:" + v_type + " " + "VALUE:" + v_value + "\n")
}
func gernerate_Opt(file string, output *os.File) {
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := NewParser(br)
	p.SplitToken()
	p.Show()
	expr := NewExpr()
	name := NewNameTable()
	v_type, v_value := expr.GetOpt(p.Content(), name)
	//fmt.Println("TYPE : " + v_type)
	//fmt.Println("VALUE : " + v_value)
	output.WriteString("TYPE:" + v_type + " " + "VALUE:" + v_value + "\n")
}
func gernerate_Value(file string, output *os.File) {
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := NewParser(br)
	p.SplitToken()
	p.Show()
	expr := NewExpr()
	name := NewNameTable()
	v_type, v_value := expr.GetValue(p.Content(), name)
	//fmt.Println("TYPE : " + v_type)
	//fmt.Println("VALUE : " + v_value)
	output.WriteString("TYPE:" + v_type + " " + "VALUE:" + v_value + "\n")
}
func Test_GetFactor(t *testing.T) {
	count := 6
	for i := 0; i < count; i++ {

		inputfile := "TestCases_Expr/Factor" + strconv.FormatInt(int64(i), 10)
		outputfile, err2 := os.Create("TestOutputs_Expr/Factor" + strconv.FormatInt(int64(i), 10))
		if err2 == nil {
			defer outputfile.Close()
		}
		if err2 != nil {
			t.Error("the file can't be opened\n")
			return
		}
		gernerate_Factor(inputfile, outputfile)
		check_output := "TestOutputs_Expr/Factor" + strconv.FormatInt(int64(i), 10)
		check_result := "TestResults_Expr/Factor" + strconv.FormatInt(int64(i), 10)
		code, match := tools.CompareTwoFiles(check_output, check_result, false)
		switch code {
		case 0:
			if match == true {
				t.Log("the test of GetFactor function passes!")
			} else {
				t.Error("the test of GetFactor function doesn't pass")
			}
		case 1:
			t.Error("there are some problems while opening the file to check")
		case 2:
			t.Error("there are some problems while loading the file to check")
		default:
			t.Error("unknown errors")
		}

	}
}
func Test_GetTerm(t *testing.T) {
	count := 8
	for i := 0; i < count; i++ {

		inputfile := "TestCases_Expr/Term" + strconv.FormatInt(int64(i), 10)
		outputfile, err2 := os.Create("TestOutputs_Expr/Term" + strconv.FormatInt(int64(i), 10))
		if err2 == nil {
			defer outputfile.Close()
		}
		if err2 != nil {
			t.Error("the file can't be opened\n")
			return
		}
		gernerate_Term(inputfile, outputfile)
		check_output := "TestOutputs_Expr/Term" + strconv.FormatInt(int64(i), 10)
		check_result := "TestResults_Expr/Term" + strconv.FormatInt(int64(i), 10)
		code, match := tools.CompareTwoFiles(check_output, check_result, false)
		switch code {
		case 0:
			if match == true {
				t.Log("the test of GetTerm function passes!")
			} else {
				t.Error("the test of GetTerm function doesn't pass")
			}
		case 1:
			t.Error("there are some problems while opening the file to check")
		case 2:
			t.Error("there are some problems while loading the file to check")
		default:
			t.Error("unknown errors")
		}

	}
}
func Test_GetOpt(t *testing.T) {
	count := 7
	for i := 0; i < count; i++ {

		inputfile := "TestCases_Expr/Opt" + strconv.FormatInt(int64(i), 10)
		outputfile, err2 := os.Create("TestOutputs_Expr/Opt" + strconv.FormatInt(int64(i), 10))
		if err2 == nil {
			defer outputfile.Close()
		}
		if err2 != nil {
			t.Error("the file can't be opened\n")
			return
		}
		gernerate_Opt(inputfile, outputfile)
		check_output := "TestOutputs_Expr/Opt" + strconv.FormatInt(int64(i), 10)
		check_result := "TestResults_Expr/Opt" + strconv.FormatInt(int64(i), 10)
		code, match := tools.CompareTwoFiles(check_output, check_result, false)
		switch code {
		case 0:
			if match == true {
				t.Log("the test of GetOpt function passes!")
			} else {
				t.Error("the test of GetOpt function doesn't pass")
			}
		case 1:
			t.Error("there are some problems while opening the file to check")
		case 2:
			t.Error("there are some problems while loading the file to check")
		default:
			t.Error("unknown errors")
		}

	}
}
func Test_GetExpr(t *testing.T) {
	count := 8
	for i := 0; i < count; i++ {

		inputfile := "TestCases_Expr/Expr" + strconv.FormatInt(int64(i), 10)
		outputfile, err2 := os.Create("TestOutputs_Expr/Expr" + strconv.FormatInt(int64(i), 10))
		if err2 == nil {
			defer outputfile.Close()
		}
		if err2 != nil {
			t.Error("the file can't be opened\n")
			return
		}
		gernerate_Expr(inputfile, outputfile)
		check_output := "TestOutputs_Expr/Expr" + strconv.FormatInt(int64(i), 10)
		check_result := "TestResults_Expr/Expr" + strconv.FormatInt(int64(i), 10)
		code, match := tools.CompareTwoFiles(check_output, check_result, false)
		switch code {
		case 0:
			if match == true {
				t.Log("the test of GetExpr function passes!")
			} else {
				t.Error("the test of GetExpr function doesn't pass")
			}
		case 1:
			t.Error("there are some problems while opening the file to check")
		case 2:
			t.Error("there are some problems while loading the file to check")
		default:
			t.Error("unknown errors")
		}

	}
}

func Test_GetValue(t *testing.T) {
	count := 5
	for i := 0; i < count; i++ {

		inputfile := "TestCases_Expr/Value" + strconv.FormatInt(int64(i), 10)
		outputfile, err2 := os.Create("TestOutputs_Expr/Value" + strconv.FormatInt(int64(i), 10))
		if err2 == nil {
			defer outputfile.Close()
		}
		if err2 != nil {
			t.Error("the file can't be opened\n")
			return
		}
		gernerate_Value(inputfile, outputfile)
		check_output := "TestOutputs_Expr/Value" + strconv.FormatInt(int64(i), 10)
		check_result := "TestResults_Expr/Value" + strconv.FormatInt(int64(i), 10)
		code, match := tools.CompareTwoFiles(check_output, check_result, false)
		switch code {
		case 0:
			if match == true {
				t.Log("the test of GetValue function passes!")
			} else {
				t.Error("the test of GetValue function doesn't pass")
			}
		case 1:
			t.Error("there are some problems while opening the file to check")
		case 2:
			t.Error("there are some problems while loading the file to check")
		default:
			t.Error("unknown errors")
		}
	}
}

