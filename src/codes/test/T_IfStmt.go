package main

import (
	"bufio"
	"codes/interpreter"
	"fmt"
	"os"
)

func testParser() {
	file := "test_parser_ifstmt"
	fin, err := os.Open(file)
	br := bufio.NewReader(fin)
	defer fin.Close()
	if err != nil {
		fmt.Println("Error while loading file!")
		return
	}
	p := interpreter.NewParser(br)
	p.SplitToken()
	if_stmt := interpreter.NewIfStmt()
	name := interpreter.NewNameTable()
	if_stmt.Execute(p.Content(), name)
}

func main() {
	testParser()
}
