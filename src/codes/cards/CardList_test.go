package cards

/***************************************************************************************
**文件名：CardList_test.go
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
	"strings"
	"testing"
)

func Test_CardList_AddCard(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_CardList/AddCard")
	outputfile, err2 := os.Create("TestOutputs_CardList/AddCard")
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
			t.Error("there are some problems while loading file TestCases_CardList/AddCard\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		testlist := NewCardList()
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) >= 2 {
				toadd := NewCard(stringslices[i])
				testlist.AddCard(*toadd)
			}
		}
		for i := 0; i < testlist.Length; i++ {
			outputfile.WriteString(testlist.Cards[i].ToString() + " ")
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_CardList/AddCard", "TestResults_CardList/AddCard", false)
	switch code {
	case 0:
		if match == true {
			t.Log("the test of AddCard function passes!")
		} else {
			t.Error("the test of AddCard function doesn't pass")
		}
	case 1:
		t.Error("there are some problems while opening the file to check")
	case 2:
		t.Error("there are some problems while loading the file to check")
	default:
		t.Error("unknown errors")
	}
}

func Test_CardList_Standardize(t *testing.T) {
	casesfile, err1 := os.Open("TestCases_CardList/Standardize")
	outputfile, err2 := os.Create("TestOutputs_CardList/Standardize")
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
			t.Error("there are some problems while loading file TestCases_CardList/Standardize\n")
			return
		}
		stringslices := strings.Split(string(buf), " ")
		testlist := NewCardList()
		for i := 0; i < len(stringslices); i++ {
			if len(stringslices[i]) >= 2 {
				toadd := NewCard(stringslices[i])
				testlist.AddCard(*toadd)
			}
		}
		testlist.Standardize()
		for i := 0; i < testlist.Length; i++ {
			outputfile.WriteString(testlist.Cards[i].ToString() + " ")
		}
		outputfile.WriteString("\n")
	}
	code, match := tools.CompareTwoFiles("TestOutputs_CardList/Standardize", "TestResults_CardList/Standardize", false)
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

