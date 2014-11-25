package cards

/***************************************************************************************
**文件名：Card_test.go
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

//测试Int2Colour函数
func Test_Card_Int2Colour(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/Int2Colour")
	outputfile, err2 := os.Create("TestOutputs_Card/Int2Colour")
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
			t.Error("there are some problems while loading file TestCases_Card/Int2Colour\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) == 0 {
				continue
			}
			num, error := strconv.ParseInt(stringslices[i], 10, 32)
			if error != nil {
				t.Error("illegal number form " + stringslices[i] + "\n")
				return
			} else {
				symbol := Int2Colour(int(num))
				outputfile.WriteString(symbol + " ")
			}
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/Int2Colour", "TestResults_Card/Int2Colour", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of Int2Colour function passes!")
		} else {
			t.Error("the test of Int2Colour function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

//测试Int2Point函数
func Test_Card_Int2Point(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/Int2Point")
	outputfile, err2 := os.Create("TestOutputs_Card/Int2Point")
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
			t.Error("there are some problems while loading file TestCases_Card/Int2Point\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) == 0 {
				continue
			}
			num, error := strconv.ParseInt(stringslices[i], 10, 32)
			if error != nil {
				t.Error("illegal number form " + stringslices[i] + "\n")
				return
			} else {
				symbol := Int2Point(int(num))
				outputfile.WriteString(symbol + " ")
			}
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/Int2Point", "TestResults_Card/Int2Point", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of Int2Point function passes!")
		} else {
			t.Error("the test of Int2Point function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

//测试ColourIncrease函数
func Test_Card_ColourIncrease(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/ColourIncrease")
	outputfile, err2 := os.Create("TestOutputs_Card/ColourIncrease")
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
			t.Error("there are some problems while loading file TestCases_Card/ColourIncrease\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) == 0 {
				continue
			}
			testcard := NewCard(stringslices[i])
			if testcard == nil {
				outputfile.WriteString("XX ")
			} else {
				testcard.Colour = ColourIncrease(testcard.Colour)
				outputfile.WriteString(testcard.ToString() + " ")
			}
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/ColourIncrease", "TestResults_Card/ColourIncrease", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of ColourIncrease function passes!")
		} else {
			t.Error("the test of ColourIncrease function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

//测试PointIncrease函数
func Test_Card_PointIncrease(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/PointIncrease")
	outputfile, err2 := os.Create("TestOutputs_Card/PointIncrease")
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
			t.Error("there are some problems while loading file TestCases_Card/PointIncrease\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) == 0 {
				continue
			}
			testcard := NewCard(stringslices[i])
			if testcard == nil {
				outputfile.WriteString("XX ")
			} else {
				testcard.Point = PointIncrease(testcard.Point)
				outputfile.WriteString(testcard.ToString() + " ")
			}
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/PointIncrease", "TestResults_Card/PointIncrease", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of PointIncrease function passes!")
		} else {
			t.Error("the test of PointIncrease function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

//测试GetPointIndex函数
func Test_Card_GetPointIndex(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/GetPointIndex")
	outputfile, err2 := os.Create("TestOutputs_Card/GetPointIndex")
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
			t.Error("there are some problems while loading file TestCases_Card/GetPointIndex\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) == 0 {
				continue
			}
			testcard := NewCard(stringslices[i])
			if testcard == nil {
				outputfile.WriteString("X ")
			} else {
				pointindex := testcard.GetPointIndex()
				outputfile.WriteString(strconv.FormatInt(int64(pointindex), 10) + " ")
			}
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/GetPointIndex", "TestResults_Card/GetPointIndex", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of GetPointIndex function passes!")
		} else {
			t.Error("the test of GetPointIndex function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

//测试GetColourIndex函数
func Test_Card_GetColourIndex(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/GetColourIndex")
	outputfile, err2 := os.Create("TestOutputs_Card/GetColourIndex")
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
			t.Error("there are some problems while loading file TestCases_Card/GetColourIndex\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) == 0 {
				continue
			}
			testcard := NewCard(stringslices[i])
			if testcard == nil {
				outputfile.WriteString("X ")
			} else {
				colourindex := testcard.GetColourIndex()
				outputfile.WriteString(strconv.FormatInt(int64(colourindex), 10) + " ")
			}
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/GetColourIndex", "TestResults_Card/GetColourIndex", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of GetColourIndex function passes!")
		} else {
			t.Error("the test of GetColourIndex function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

//测试NewCardByIndex函数
func Test_Card_NewCardByIndex(t *testing.T) {
	casesfile, err := os.Open("TestCases_Card/NewCardByIndex")
	if err != nil {
		t.Error("the file can't be opened\n")
		return
	}
	defer casesfile.Close()
	reader := bufio.NewReader(casesfile)
	for {
		buf, _, error := reader.ReadLine()
		if error == io.EOF {
			break
		} else if error != nil {
			t.Error("there are some problems while loading file TestCases_Card/GetColourIndex\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		if len(stringslices) == 2 {
			colourindex, err1 := strconv.ParseInt(stringslices[0], 10, 32)
			pointindex, err2 := strconv.ParseInt(stringslices[1], 10, 32)
			if err1 != nil || err2 != nil {
				t.Error("illegal number format\n")
				return
			}
			testcard := NewCardByIndex(int(colourindex), int(pointindex))
			retcolour := testcard.GetColourIndex()
			retpoint := testcard.GetPointIndex()
			if retcolour != int(colourindex) || retpoint != int(pointindex) {
				t.Error("the test of NewCardByIndex function doesn't pass")
				return
			}
		}
	}
	t.Log("the test of NewCardByIndex function passes!")
	return
}

//测试BreakRegEx函数
func Test_Card_BreakRegEx(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/BreakRegEx")
	outputfile, err2 := os.Create("TestOutputs_Card/BreakRegEx")
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
			t.Error("there are some problems while loading file TestCases_Card/GetColourIndex\n")
			return
		}
		colour, point, ok := BreakRegEx(string(buf))
		if ok == true {
			outputfile.WriteString(colour + " " + point + "\n")
		} else {
			outputfile.WriteString("false\n")
		}
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/BreakRegEx", "TestResults_Card/BreakRegEx", true)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of BreakRegEx function passes!")
		} else {
			t.Error("the test of BreakRegEx function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

func Test_Card_Match(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_Card/Match")
	outputfile, err2 := os.Create("TestOutputs_Card/Match")
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
			t.Error("there are some problems while loading file TestCases_Card/GetColourIndex\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		if len(stringslices) >= 3 && len(stringslices[0]) != 0 && len(stringslices[2]) != 0 {
			testcard := NewCard(stringslices[0])
			if testcard == nil {
				outputfile.WriteString("\n")
				continue
			}
			fit := testcard.Match(stringslices[1], stringslices[2])
			outputfile.WriteString(strconv.FormatBool(fit) + "\n")
		} else {
			outputfile.WriteString("\n")
		}
	}
	code, match := tools.CompareTwoFiles("TestOutputs_Card/Match", "TestResults_Card/Match", true)
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

