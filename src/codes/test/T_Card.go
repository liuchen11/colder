package main

import(
	"fmt"
	"bufio"
	"os"
	"codes/cards"
	"strings"
)

func TestCardBase(list []cards.Card) {
	for i:=0;i<len(list);i++ {
		list[i].Print()
		fmt.Print(" ",list[i].GetColourIndex()," ",list[i].GetPointIndex(),"\n")
	}
}

func TestCardMatch(list []cards.Card,regexs []string) {
	for j:=0;j<len(regexs);j++{
		fmt.Println("Match ",regexs[j])
		for i:=0;i<len(list);i++ {
			list[i].Print()
			fmt.Print(" ",list[i].Match(regexs[j])," ",regexs[j],"\n")
		}
	}
}

func main() {
	if len(os.Args)<3{
		fmt.Println("the number of arguments must be at least 3")
		return
	}
	file,err:=os.Open(os.Args[1])
	defer file.Close()
	if err!=nil{
		fmt.Println("can't open the file ",os.Args[1])
		return
	}
	list:=make([]cards.Card,0,10)
	reader:=bufio.NewReader(file)
	if reader==nil{
		fmt.Println("can't read from file ",os.Args[1])
		return
	}
	buf,_,error:=reader.ReadLine()
	for ;error==nil;{
		buffer:=strings.Split(string(buf)," ")
		//fmt.Println(buffer)
		for i:=0;i<len(buffer);i++{
			//fmt.Println(buffer[i])
			if len(buffer[i])>=2 {
				toadd:=cards.NewCard(buffer[i][0:2])
				list=append(list,*toadd)
			}
		}
		buf,_,error=reader.ReadLine()
	}
	switch os.Args[2]{
	case "base":
		TestCardBase(list)
	case "match":
		if len(os.Args)>3{
			regexfile,_:=os.Open(os.Args[3])
			defer regexfile.Close()
			regexreader:=bufio.NewReader(regexfile)
			regexs:=make([]string,0,10)
			buf,_,error=regexreader.ReadLine()
			for ;error==nil;{
				buffer:=strings.Split(string(buf)," ")
				for i:=0;i<len(buffer);i++{
					regexs=append(regexs,buffer[i])
				}
				buf,_,error=regexreader.ReadLine()
			}
			//fmt.Println(regexs)
			TestCardMatch(list,regexs)
		}else{
			fmt.Println("the number of arguments is at least 4")
		}
	default:
	}
	return
}