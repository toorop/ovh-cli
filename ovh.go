package main

import (
	//"flag"
	"bufio"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"github.com/wsxiaoys/terminal"
	"os"
	"runtime"
	"strings"
)

const (
	OVH_APP_KEY    = "SECRET"
	OVH_APP_SECRET = "SECRET"
	NL             = "\r\n"
	TAB            = "\t"
)

var (
	// ck represents the cocumer key
	ck string
)

func init() {
	// Consumer key
	ck = os.Getenv("OVH_CONSUMER_KEY")
	fmt.Println(ck)

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

func main() {
	app := cli.NewApp()
	app.Name = "ovh"
	app.Usage = "ovh is a command line interface which interact with OVH services using their API."
	app.Version = "2.0"
	app.Author = "Stéphane Depierrepont aka Toorop"
	app.Email = "toorop@toorop.fr"
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [section] [subcommand] [arguments]

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
	app.Action = func(c *cli.Context) {
		println("Hello friend!")
	}

	// New govh client
	client := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)

	// IP
	serverCmd := []cli.Command{
		{
			Name:        "list",
			Usage:       "List your IPS",
			Description: "You can add a IP type as argument. Example: ovh ip list dedicated",
			Action: func(c *cli.Context) {
				var resp string
				// New ip ressource
				ipr, err := ip.New(client)
				if err != nil {
					return
				}
				ipType := "all"
				if c.Args().First() != "" {
					ipType = strings.ToLower(c.Args().First())
				}
				ips, err := ipr.List(ipType)
				if err != nil {
					dieError(err)
				}
				for _, i := range ips {
					resp = fmt.Sprintf("%s%s\r\n", resp, i.IP)
				}
				if len(resp) > 2 {
					resp = resp[0 : len(resp)-2]
				}
				dieOk(resp)
			},
		},
		{
			Name:  "remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) {
				println("removed task template: ", c.Args().First())
			},
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "ip",
			Usage:       "Define IP section",
			Subcommands: serverCmd,
		},
	}

	app.Run(os.Args)
}
