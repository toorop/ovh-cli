package main

import (
	//"flag"
	"bufio"
	"fmt"

	"github.com/toorop/govh"
	//"github.com/Toorop/govh/ip"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"github.com/wsxiaoys/terminal"
	//"strings"
)

const (
	NL      = "\r\n"
	TAB     = "   "
	NLTAB   = NL + TAB
	VERSION = "2.0.1"
)

var (
	// ck represents the consumer key
	ck string
	// region represents API region.
	// EU (default) or CA
	region string
)

func init() {
	// region
	region = os.Getenv("OVH_REGION")

	// Consumer key
	ck = os.Getenv("OVH_CONSUMER_KEY")

	// if No ConsumerKey, request one
	if len(ck) == 0 {
		var r []byte

		if runtime.GOOS == "windows" {
			fmt.Println(NL, "No consumer key found in environment vars !", NL)
		} else {
			terminal.Stdout.Clear().Move(0, 0).Color("r").
				Print("No consumer key found in environment vars !").Nl().Nl().Reset()
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
			fmt.Println("and restart ovh CLI application.", NL)
			os.Exit(0)
		}

		ck, link, err := govh.AuthGetConsumerKey(getAppKey(region), region)
		if err != nil {
			dieError(err)
		}
		fmt.Print("\r\nYour consumer key is : ")
		if runtime.GOOS != "windows" {
			terminal.Stdout.Color("g").Print(ck).Nl().Reset().Nl()
		} else {
			fmt.Print(ck)
		}

		fmt.Println("Now you need to validate it :")
		if runtime.GOOS != "windows" {
			fmt.Printf("\t- If you have a browser available on this machine it will open to the validation page.\n\t- If not, copy and paste the link below in a browser to validate your key :\r\n\r\n%s\r\n", link)
			webbrowser.Open(link)
		} else {
			fmt.Printf("To do it just copy and paste the link below in a browser and follow instructions on OVH website :\r\n\r\n%s\r\n", link)
		}

		fmt.Println("\r\nWhen it will be done run the following command :")
		if runtime.GOOS == "windows" {
			fmt.Printf("SET OVH_CONSUMER_KEY=%s%s%s", ck, NL, NL)
		} else {
			fmt.Printf("export OVH_CONSUMER_KEY=%s%s%s", ck, NL, NL)
		}
		fmt.Println("and restart ovh CLI application.")
		os.Exit(0)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "ovh"
	app.Usage = "ovh-cli brings OVH services to the command line."
	app.Version = VERSION
	app.Author = "Stèphane Depierrepont aka Toorop"
	app.Email = "toorop@toorop.fr"
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

SECTIONS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Description}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.CommandHelpTemplate = `
   {{.Name}} - {{.Description}}

USAGE:
   {{.Usage}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.SubcommandHelpTemplate = `
   {{.Name}} - {{.Usage}}

SUBS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Description}}
   {{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	// New govh client
	client := govh.New(getAppKey(region), getAppSecret(region), ck, region)

	// default action: help
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	// Main  getFwCmds
	app.Commands = []cli.Command{
		{
			Name:        "me",
			Usage:       "me subsection",
			Description: "Commands about me",
			Subcommands: getMeCmds(client),
		},
		{
			Name:        "domain",
			Usage:       "domain subsection",
			Description: "Commands about domains",
			Subcommands: getDomainCmds(client),
		},
		{
			Name:        "ip",
			Usage:       "IP subsection",
			Description: "Commands about IP",
			Subcommands: getIPCmds(client),
		}, {
			Name:        "fw",
			Usage:       "Firewall subsection",
			Description: "Commands OVH firewall",
			//	Subcommands: getFwCmds(client),
		}, {
			Name:        "server",
			Usage:       "Server subsection",
			Description: "Commands about OVH server",
			Subcommands: getServerCmds(client),
		}, {
			Name:        "sms",
			Usage:       "Sms subsection",
			Description: "Commands about OVH SMS",
			Subcommands: getSmsCmds(client),
		}, {
			Name:        "spam",
			Usage:       "Spam subsection",
			Description: "Commands about OVH antispam protection",
			Subcommands: getSpamCmds(client),
		}, {
			Name:        "cloud",
			Usage:       "Cloud subsection",
			Description: "Commands about OVH cloud",
			Subcommands: getCloudCmds(client),
		}, {
			Name:        "dedicatedcloud",
			Usage:       "Dedicated Cloud subsection",
			Description: "Commands about OVH Dedicated Cloud",
			Subcommands: getDedicatedCloudCmds(client),
		},
	}

	app.Run(os.Args)
}
