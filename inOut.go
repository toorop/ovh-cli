package main

import (
	"fmt"
	"github.com/wsxiaoys/terminal"
	"os"
	"runtime"
)

func dieError(v ...interface{}) {
	if runtime.GOOS == "windows" {
		fmt.Println("Error : ", v, NL)
	} else {
		terminal.Stdout.Color("r").Print("Error : ", v).Nl().Reset()
	}
	os.Exit(1)
}

func dieOk(r string) {
	fmt.Println(r)
	os.Exit(0)
}

func debug(v ...interface{}) {
	terminal.Stdout.Color("y").Print("Debug : ", v).Nl().Reset()
}
