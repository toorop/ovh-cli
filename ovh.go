package main

import (
	"flag"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/toqueteos/webbrowser"
	"github.com/wsxiaoys/terminal"
	//"github.com/wsxiaoys/terminal/color"
	"bufio"
	//"encoding/json"
	"os"
	"runtime"
)

const (
	OVH_APP_KEY    = "SECRET"
	OVH_APP_SECRET = "SECRET"
	NL             = "\r\n"
)

var (
	ck  string // consumer key
	cmd Cmd
)

func init() {
	flag.StringVar(&ck, "ck", "", "Consumer Key")
	//flag.StringVar(&outputFormat, "of", "JSON", "Output format")
	flag.Parse()

	if len(flag.Args()) > 0 {
		cmd = Cmd{
			Domain: flag.Arg(0),
			Action: flag.Arg(1),
			Args:   flag.Args(),
		}
	}

	// WYAUsR31Z3dT9Y5f0arTHeZwpFRdcnz2
	ck = os.Getenv("OVH_CONSUMER_KEY")

	// if No ConsumerKey, request one
	if len(ck) == 0 {
		var r []byte

		if runtime.GOOS == "windows" {
			fmt.Println(NL, "No consumer key found in environnement vars !", NL)
		} else {
			terminal.Stdout.Clear().Move(0, 0).Color("r").
				Print("No consumer key found in environnement vars !").Nl().Nl().Reset()
		}
		for {
			fmt.Print("Have you a valid Consumer Key for that app ? (y/n) : ")

			r, _, _ = bufio.NewReader(os.Stdin).ReadLine()
			if r[0] == 110 || r[0] == 121 {
				break
			}
		}
		// Yes
		if r[0] == 121 {
			fmt.Println("\r\nRun the following command :", NL)
			if runtime.GOOS == "windows" {
				fmt.Println("SET OVH_CONSUMER_KEY=your_consumer_key", NL)
			} else {
				fmt.Println("export OVH_CONSUMER_KEY=your_consumer_key", NL)
			}
			fmt.Println("and restart ovh CLI application.\r\n")
			os.Exit(0)
		}

		ck, link, err := govh.AuthGetConsumerKey(OVH_APP_KEY)
		if err != nil {
			panic(err)
		}
		fmt.Print("\r\nYour consumer key is : ")
		if runtime.GOOS != "windows" {
			terminal.Stdout.Color("g").Print(ck).Nl().Reset().Nl()
		} else {
			fmt.Print(ck)
		}

		fmt.Println("Now you need to validate it :")
		if runtime.GOOS != "windows" {
			fmt.Printf("\t- If you have a browser available on this machine it will open to the validation page.\n\t- If not copy and paste the link below in a browser to validate your key :\r\n\r\n%s\r\n", link)
			webbrowser.Open(link)
		} else {
			fmt.Printf("To do it just copy and paste the link below in a browser and follow instructions on OVH website :\r\n\r\n%s\r\n", link)
		}

		fmt.Println("\r\nWhen it will be done run the following command : \r\n")
		if runtime.GOOS == "windows" {
			fmt.Printf("SET OVH_CONSUMER_KEY=%s\r\n", ck)
		} else {
			fmt.Printf("export OVH_CONSUMER_KEY=%s\r\n", ck)
		}

		fmt.Println("and restart ovh CLI application.\r\n")
		os.Exit(0)
	}
}

// Main
func main() {
	var resp string
	var err error

	switch cmd.Domain {
	case "ip":
		resp, err = ipHandler(&cmd)
		break
	case "help":
		resp = "See : https://github.com/Toorop/govh"
		break
	default:
		dieError("This section '", cmd.Domain, "' is not valid or not implemented yet !")
	}
	if err != nil {
		dieError(err, resp)
	}
	dieOk(resp)

}
