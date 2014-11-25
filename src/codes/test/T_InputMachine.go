package main

import (
	"bufio"
	"codes/interpreter"
	"fmt"
	"os"
)

func testGetNextChar() {
	file := "justfortest"
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("File Input Error!")
		return
	}
	Machine := interpreter.NewInputMachine(br)
	for {
		_, err := Machine.GetNextChar()
		if err != nil {
			break
		}
	}
}

func testMatchNextToken() {
	file := "justfortest"
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	Machine := interpreter.NewInputMachine(br)
	for {
		if Machine.Endflag == true {
			break
		}
		Machine.MatchNextToken()
	}
}

func main() {
	switch os.Args[1] {
	case "char":
		testGetNextChar()
	case "token":
		testMatchNextToken()
	case "func":
		if len(os.Args) >= 3 {
			tokens := interpreter.Analyse(os.Args[2])
			fmt.Println(tokens)
		}
	default:
	}
}
