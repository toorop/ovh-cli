package main

import (
	"fmt"
	"github.com/wsxiaoys/terminal"
	"os"
	//"runtime"
)

func dieError(v ...interface{}) {
	//if runtime.GOOS == "windows" {
	fmt.Println("ERROR!")
	for _, vv := range v {
		fmt.Printf("%v", vv)
	}
	fmt.Println("")
	/*} else {
		terminal.Stdout.Color("r").Print("Error : ")
		for _, vv := range v {
			terminal.Stdout.Print(vv)
		}
		terminal.Stdout.Nl().Reset()
	}*/
	os.Exit(1)
}

// dieInvalidKey will exit in case of client key has expired or
// is not valid
func dieInvalidConsumerKey() {
	dieError("Your credentials seems to have expired.", NL, "Delete environement variable OVH_CONSUMER_KEY and relaunch ovh-cli to generate a new one.", NL, "On Linux|MacOS: export OVH_CONSUMER_KEY=", NL, "On windows: SET OVH_CONSUMER_KEY=")
}

// Exit & and display error on bad arguments
func dieBadArgs() {
	dieError("Bad arg(s). Run ./ovh command [subCommand...] --help for help.")
}

// Exit if args are missing
func dieIfArgsMiss(nbArgs, requiered int) {
	if nbArgs < requiered {
		dieBadArgs()
	}
}

func dieOk(r ...string) {
	if len(r) != 0 {
		for _, line := range r {
			fmt.Println(line)
		}
	}
	os.Exit(0)
}

func dieDone() {
	dieOk("Done!")
}

func debug(v ...interface{}) {
	terminal.Stdout.Color("y").Print("Debug : ", v).Nl().Reset()
}
