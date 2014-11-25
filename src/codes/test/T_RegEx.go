package main

import(
	"os"
	"bufio"
	"fmt"
	"codes/cards"
)

func TestStd(input string){
	test:=cards.NewRegEx(input)
	test.Stdize()
	fmt.Println(test.Value)
}

func TestCompile(input string){
	test:=cards.NewRegEx(input)
	test.Stdize()
	fmt.Println(test.Value)
	test.Compile()
	parts:=test.GetParts()
	for i:=0;i<len(parts);i++{
		parts[i].Print()
	}
}

func main(){
	if len(os.Args)>1{
		switch os.Args[1]{
			case "std":
				reader:=bufio.NewReader(os.Stdin)
				input,_,_:=reader.ReadLine()
				TestStd(string(input))
			case "compile":
				reader:=bufio.NewReader(os.Stdin)
				input,_,_:=reader.ReadLine()
				TestCompile(string(input))
			default:
		}
	}
	
}