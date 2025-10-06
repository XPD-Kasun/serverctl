package main

import (
	"cli/core"
	"cli/providers/vsftp"
	"fmt"
	"runtime"
)

func main() {

	os := runtime.GOOS
	fmt.Println(core.File, os, core.DetectOS())

	fmt.Println(core.DetectPkgManager())
	// var s core.Provider
	// var a A
	// a.Fn = 34
	// var b B
	// b.A = a
	// b.test2()
	// b.test()
	// str, err := xml.MarshalIndent(runtime.MemStats{}, "", "  ")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Println(string(str))

	var provider core.Provider = vsftp.NewVSFTPProvider()
	isExists, err := provider.Exists()
	if !isExists {
		println("not found" + err.Error())
		return
	}
	println("found" + err.Error())
}

type A struct {
	Fn int `json:"price"`
}

func (a A) test() {
	fmt.Printf("a.a: %v\n", a.Fn)
}

type B struct {
	A
	q int
}

func (b B) test2() {
	b.test()
	fmt.Printf("b.q: %v\n", b.q)
}
