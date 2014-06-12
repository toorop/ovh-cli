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
	//"strings"
)

const (
	NL      = "\r\n"
	TAB     = "   "
	NLTAB   = NL + TAB
	VERSION = "2.0"
)

var (
	// ck represents the cocumer key
	ck string
	// region represents API region.
	// EU (default) or CA
	region string
	// err
	err error
	// IP Ressource
	ipr *ip.IpRessource
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

		fmt.Println("\r\nWhen it will be done run the following command : \r\n")
		if runtime.GOOS == "windows" {
			fmt.Printf("SET OVH_CONSUMER_KEY=%s%s%s", ck, NL, NL)
		} else {
			fmt.Printf("export OVH_CONSUMER_KEY=%s%s%s", ck, NL, NL)
		}
		fmt.Println("and restart ovh CLI application.\r\n")
		os.Exit(0)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "ovh"
	app.Usage = "ovh is a command line interface which interact with OVH API."
	app.Version = VERSION
	app.Author = "Stéphane Depierrepont aka Toorop"
	app.Email = "toorop@toorop.fr"
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [section] [subsection...] [command] [arguments]

SECTIONS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.CommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Description}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.SubcommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [subsection] command [command options] [arguments...]

COMMANDS|SUBSECTION:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	// New govh client
	client := govh.NewClient(getAppKey(region), getAppSecret(region), ck, region)

	// default action: help
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	// Main  getFwCmds
	app.Commands = []cli.Command{
		{
			Name:        "ip",
			Usage:       "IP section",
			Description: "IP commands",
			Subcommands: getIpCmds(client),
		}, {
			Name:        "fw",
			Usage:       "Firewall subsection",
			Description: "Firewall commands",
			Subcommands: getFwCmds(client),
		},
	}

	app.Run(os.Args)
}
