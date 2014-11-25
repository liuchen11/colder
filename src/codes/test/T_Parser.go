package main

import (
	"bufio"
	"codes/interpreter"
	"fmt"
	"os"
)

func testParser() {
	file := "test_expr"
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := interpreter.NewParser(br)
	p.SplitToken()
	p.Show()
}

func main() {
	testParser()
}
