package main

import (
	"cli/core"
	"fmt"
	"runtime"
)

/*
 */
func main() {
	goos := runtime.GOOS
	fmt.Println(core.File, goos)
}
