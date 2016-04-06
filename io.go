package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	//"runtime/debug"
)

// dieError outputs error and exit 1
func dieError(err error) {
	fmt.Println("Error:", err)
	/*if len(v) != 0 {
		fmt.Println(v[0])
	}*/
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
	dieError(fmt.Errorf("Your credentials seems to have expired." + "Delete environement variable OVH_CONSUMER_KEY and relaunch ovh-cli to generate a new one." + "On Linux|MacOS: export OVH_CONSUMER_KEY=" + "On windows: SET OVH_CONSUMER_KEY="))
}

// Exit & and display error on bad arguments
func dieBadArgs(msg ...string) {
	errMsg := "Bad arg(s). Run ./ovh command [subCommand...] --help for help."
	if len(msg) > 0 {
		errMsg = msg[0]
	}
	dieError(fmt.Errorf(errMsg))
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
	out := ""

	// slice ?
	s := reflect.ValueOf(data)
	switch s.Kind() {
	case reflect.Slice:
		ret := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}

		for _, v := range ret {
			out += fmt.Sprintf("%v\n", v)
		}
	default:
		out = fmt.Sprintf("%v", data)
	}

	// clean ending
	if strings.HasSuffix(out, "\n") {
		out = out[:len(out)-1]
	}
	return out
}

/*func debug(v ...interface{}) {
	terminal.Stdout.Color("y").Print("Debug : ", v).Nl().Reset()
}*/
