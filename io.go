package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"runtime/debug"
)

// dieError outputs error and exit 1
func dieError(v ...interface{}) {
	fmt.Print("ERR")
	if len(v) != 0 {
		fmt.Println(v)
	}
	os.Exit(1)
}

// dieOnError check err and die on err
func dieOnError(err error) {
	if err == nil {
		return
	}
	dieError(err)
}

// dieInvalidKey will exit in case of client key has expired or
// is not valid
func dieInvalidConsumerKey() {
	dieError("Your credentials seems to have expired.", NL, "Delete environement variable OVH_CONSUMER_KEY and relaunch ovh-cli to generate a new one.", NL, "On Linux|MacOS: export OVH_CONSUMER_KEY=", NL, "On windows: SET OVH_CONSUMER_KEY=")
}

// Exit & and display error on bad arguments
func dieBadArgs(msg ...string) {
	errMsg := "Bad arg(s). Run ./ovh command [subCommand...] --help for help."
	if len(msg) > 0 {
		errMsg = msg[0]
	}
	dieError(errMsg)
}

// Exit if args are missing
func dieIfArgsMiss(nbArgs, requiered int) {
	if nbArgs < requiered {
		dieBadArgs()
	}
}

func dieOk(v ...interface{}) {
	if len(v) != 0 {
		fmt.Printf("%v\n", v[0])
	}
	os.Exit(0)
}

func dieDone() {
	dieOk("Done!")
}

// formatOutput return formated structure (json or raw string)
func formatOutput(data interface{}, toJSON bool) string {
	if toJSON {
		buf, err := json.Marshal(data)
		dieOnError(err)
		return string(buf)
	}
	return fmt.Sprintf("%s", data)
}

/*func debug(v ...interface{}) {
	terminal.Stdout.Color("y").Print("Debug : ", v).Nl().Reset()
}*/
