package cards

/***************************************************************************************
**文件名：RegEx_test.go
**包名称：cards
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"bufio"
	"codes/tools"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_regex_range_range_add(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_RegEx/range_add")
	outputfile, err2 := os.Create("TestOutputs_RegEx/range_add")
	if err1 == nil {
		defer casesfile.Close()
	}
	if err2 == nil {
		defer outputfile.Close()
	}
	if err1 != nil || err2 != nil {
		t.Error("the file can't be opened\n")
		return
	}
	reader := bufio.NewReader(casesfile)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error("there are some problems while loading file TestCases_RegEx/range_add\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		testrange := new(regex_range)
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) > 0 {
				testrange.add(stringslices[i])
			}
		}
		for i := 0; i < testrange.length; i++ {
			outputfile.WriteString(testrange.content[i] + " ")
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_RegEx/range_add", "TestResults_RegEx/range_add", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of range_add function passes!")
		} else {
			t.Error("the test of range_add function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

func Test_RegEx_Standardize(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_RegEx/Standardize")
	outputfile, err2 := os.Create("TestOutputs_RegEx/Standardize")
	if err1 == nil {
		defer casesfile.Close()
	}
	if err2 == nil {
		defer outputfile.Close()
	}
	if err1 != nil || err2 != nil {
		t.Error("the file can't be opened\n")
		return
	}
	reader := bufio.NewReader(casesfile)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error("there are some problems while loading file TestCases_RegEx/Standardize\n")
			return
		}
		testRegEx := NewRegEx(string(buf))
		ok := testRegEx.Standardize()
		if ok == false {
			outputfile.WriteString("false\n")
		} else {
			outputfile.WriteString(testRegEx.Value + "\n")
		}
	}
	code, match := tools.CompareTwoFiles("TestOutputs_RegEx/Standardize", "TestResults_RegEx/Standardize", true)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of Standardize function passes!")
		} else {
			t.Error("the test of Standardize function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

func Test_RegEx_Compile(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_RegEx/Compile")
	outputfile, err2 := os.Create("TestOutputs_RegEx/Compile")
	if err1 == nil {
		defer casesfile.Close()
	}
	if err2 == nil {
		defer outputfile.Close()
	}
	if err1 != nil || err2 != nil {
		t.Error("the file can't be opened\n")
		return
	}
	reader := bufio.NewReader(casesfile)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error("there are some problems while loading file TestCases_RegEx/Compile\n")
			return
		}
		testRegEx := NewRegEx(string(buf))
		testRegEx.Standardize()
		ok := testRegEx.Compile()
		if ok == true {
			for i := 0; i < len(testRegEx.parts); i++ {
				outputfile.WriteString(strconv.FormatBool(testRegEx.parts[i].fix) + " | ")
				outputfile.WriteString(strconv.FormatInt(int64(testRegEx.parts[i].times), 10) + " | ")
				for j := 0; j < testRegEx.parts[i].length; j++ {
					outputfile.WriteString(testRegEx.parts[i].content[j] + " ")
				}
				if testRegEx.parts[i].symbol != 0 {
					sym := strconv.QuoteRune(rune(testRegEx.parts[i].symbol))[1:2]
					outputfile.WriteString("| " + sym)
				}
				outputfile.WriteString("\n")
			}
		} else {
			outputfile.WriteString("false\n")
		}
		outputfile.WriteString("//\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_RegEx/Compile", "TestResults_RegEx/Compile", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of Compile function passes!")
		} else {
			t.Error("the test of Compile function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

func Test_MatchUnit(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_RegEx/MatchUnit")
	outputfile, err2 := os.Create("TestOutputs_RegEx/MatchUnit")
	if err1 == nil {
		defer casesfile.Close()
	}
	if err2 == nil {
		defer outputfile.Close()
	}
	if err1 != nil || err2 != nil {
		t.Error("the file can't be opened\n")
		return
	}
	reader := bufio.NewReader(casesfile)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error("there are some problems while loading file TestCases_RegEx/MatchUnit\n")
			return
		}
		params := strings.Split(string(buf), "||")
		if len(params) >= 4 {
			start, error := strconv.ParseInt(params[0], 10, 32)
			if error != nil {
				outputfile.WriteString("error\n")
				continue
			}
			cards := strings.Split(params[1], " ")
			cardlist := make([]Card, 0, 5)
			for i := 0; i < len(cards); i++ {
				card_to_add := NewCard(cards[i])
				if card_to_add != nil {
					cardlist = append(cardlist, *card_to_add)
				}
			}
			states := strings.Split(params[2], " ")
			if len(states) < 5 {
				outputfile.WriteString("error\n")
				continue
			}
			fix, error1 := strconv.ParseBool(states[0])
			times, error2 := strconv.ParseInt(states[1], 10, 32)
			_, error3 := strconv.ParseInt(states[2], 10, 32)
			content := strings.Split(states[3], "@")
			var symbol uint8
			if len(states[4]) != 0 {
				symbol = uint8(states[4][0])
			} else {
				symbol = 0
			}
			if error1 != nil || error2 != nil || error3 != nil {
				outputfile.WriteString("error\n")
				continue
			}
			state_range := new(regex_range)
			state_range.fix = fix
			state_range.times = int(times)
			state_range.length = len(content)
			state_range.content = content
			state_range.symbol = symbol
			vars := strings.Split(params[3], " ")
			maps := make(map[uint8]regex_var)
			for i := 0; i < len(vars); i++ {
				key_value := strings.Split(vars[i], ":")
				if len(key_value) < 2 || len(key_value[0]) == 0 {
					continue
				}
				key := key_value[0][0]
				var_params := strings.Split(key_value[1], "@")
				if len(var_params) < 3 {
					continue
				}
				var_value := var_params[0]
				var_mode, e1 := strconv.ParseInt(var_params[1], 10, 32)
				var_hasA, e2 := strconv.ParseBool(var_params[2])
				if e1 != nil || e2 != nil {
					continue
				}
				value := *NewRegex_Var(var_value, int(var_mode), var_hasA)
				maps[uint8(key)] = value
			}
			exitcode, ok := MatchUnit(int(start), cardlist, state_range, &maps)
			outputfile.WriteString(strconv.FormatInt(int64(exitcode), 10) + " " + strconv.FormatBool(ok) + "\n")
		} else {
			outputfile.WriteString("error\n")
		}
	}
	code, match := tools.CompareTwoFiles("TestOutputs_RegEx/MatchUnit", "TestResults_RegEx/MatchUnit", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of MatchUnit function passes!")
		} else {
			t.Error("the test of MatchUnit function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

func Test_Match(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_RegEx/Match")
	outputfile, err2 := os.Create("TestOutputs_RegEx/Match")
	if err1 == nil {
		defer casesfile.Close()
	}
	if err2 == nil {
		defer outputfile.Close()
	}
	if err1 != nil || err2 != nil {
		t.Error("the file can't be opened\n")
		return
	}
	reader := bufio.NewReader(casesfile)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error("there are some problems while loading file TestCases_RegEx/Match\n")
			return
		}
		parts := strings.Split(string(buf), " ")
		if len(parts) < 2 {
			outputfile.WriteString("error\n")
		} else {
			reg := NewRegEx(parts[0])
			cardlist := NewCardList()
			for i := 1; i < len(parts); i++ {
				card := NewCard(parts[i])
				if card != nil {
					cardlist.AddCard(*card)
				}
			}
			num, ok := reg.Match(*cardlist)
			outputfile.WriteString(strconv.FormatInt(int64(num), 10) + " " + strconv.FormatBool(ok) + "\n")
		}
	}
	code, match := tools.CompareTwoFiles("TestOutputs_RegEx/Match", "TestResults_RegEx/Match", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of Match function passes!")
		} else {
			t.Error("the test of Match function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

