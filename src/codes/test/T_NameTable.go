package main

import (
	"codes/interpreter"
	"fmt"
)

func showValue(ret *interpreter.NameTable) {
	fmt.Println(ret.GetString("lar0"))
	fmt.Println(ret.GetInt("lar1"))
	fmt.Println(ret.GetFloat("lar2"))
	fmt.Println(ret.GetBool("lar3"))
}
func changeValue(ret *interpreter.NameTable) {
	ret.SetString("lar0", "Largelymfs")
	ret.SetInt("lar1", 23)
	ret.SetFloat("lar2", 2.3)
	ret.SetBool("lar3", true)
	ret.SetBool("lar0", false)
	ret.SetBool("liutsen", false)
}
func main() {
	ret := interpreter.NewNameTable()
	fmt.Println("Testing....")
	fmt.Println("Phase 1 : Testing Adding Variable...")
	ret.AddVariable("string", "lar0")
	ret.AddVariable("int", "lar1")
	ret.AddVariable("float", "lar2")
	ret.AddVariable("bool", "lar3")
	ret.AddVariable("liutsen", "lar4")
	ret.AddVariable("string", "lar0")
	fmt.Println("Phase 2 : Testing Getting Variable's Type...")
	fmt.Println(ret.GetType("lar0"))
	fmt.Println(ret.GetType("lar1"))
	fmt.Println(ret.GetType("lar2"))
	fmt.Println(ret.GetType("lar3"))
	fmt.Println(ret.GetType("lar4"))
	fmt.Println(ret.GetType("lar"))
	fmt.Println("Phase 3 : Testing Changing Variable's value...")
	fmt.Println("==============Before change==============")
	showValue(ret)
	changeValue(ret)
	fmt.Println("===============After change==============")
	showValue(ret)
	ret.AddArray("stringarr", "largelymfs", 5)
	ret.SetStringArr("largelymfs", 0, "123")
	fmt.Println(ret.GetStringArr("largelymfs", 6))
	ret.AddArray("intarr", "largelymfsint", 5)
	ret.SetIntArr("largelymfsint", 1, 1)
	ret.SetIntArr("largelymfsint", 0, 0)
	ret.SetIntArr("largelymfsint", 2, 2)
	ret.SetIntArr("largelymfsint", 3, 2)
	fmt.Println(ret.GetIntArr("largelymfsint", 1))
}
