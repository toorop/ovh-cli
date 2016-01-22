package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/ip"
)

// getIpCmds return commands for Ip section
func getIPCmds(OVHClient *govh.OVHClient) (cmds []cli.Command) {
	IPClient, err := ip.New(OVHClient)
	if err != nil {
		return
	}

	// Ip commands
	cmds = []cli.Command{
		// IPBLock
		{
			Name:        "block",
			Description: "commands concerning IP block",
			Subcommands: []cli.Command{
				// List Ip Blocks
				{
					Name:        "list",
					Description: "List your IP blocks.",
					Usage:       "ovh ip block list [flag...]" + NLTAB + "Example: ovh ip block list --type vps",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "desc", Value: "", Usage: "Filter: by description (like)."},
						cli.StringFlag{Name: "ip", Value: "", Usage: "Filter: by IP (contains or equals)."},
						cli.StringFlag{Name: "routedTo", Value: "", Usage: "Filter: by routing."},
						cli.StringFlag{Name: "type", Value: "all", Usage: "Filter: by IP block type: all|cdn|dedicated|failover|hosted_ssl|housing|loadBalancing|mail|pcc|pci|private|vps|vpn|vrack|xdsl"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						fDesc := strings.ToLower(c.String("desc"))
						fIP := strings.ToLower(c.String("ip"))
						fRoutedTo := strings.ToLower(c.String("routedto"))
						fType := strings.ToLower(c.String("type"))
						if fType == "all" {
							fType = ""
						}
						IPBlocks, err := IPClient.List(fDesc, fIP, fRoutedTo, fType)
						dieOnError(err)
						// output as json ?
						if c.Bool("json") {
							buf, err := json.Marshal(IPBlocks)
							dieOnError(err)
							fmt.Println(string(buf))
						}
						for _, i := range IPBlocks {
							fmt.Println(i)
						}
						dieOk()
					},
				},
				// Get properties of a block
				{
					Name:        "properties",
					Description: "Get properties of an IP block.",
					Usage:       "ovh ip block properties IPBLOCK" + NLTAB + "Example: ovh ip block properties 91.121.228.135/32",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						block := ip.IPBlock(c.Args().First())
						properties, err := IPClient.GetBlockProperties(block)
						dieOnError(err)
						if c.Bool("json") {
							buf, err := json.Marshal(properties)
							dieOnError(err)
							fmt.Println(string(buf))
						} else {
							fmt.Println(properties.String())
						}
						dieOk()
						//dieOk(fmt.Sprintf("IP: %s%sType: %s%sDescription: %s%sRouted to: %s", properties.Ip, NL, properties.Type, NL, properties.Description, NL, properties.RoutedTo.ServiceName))
					},
				},
			}, // end of block subCommands
		}, // end of ip subcommands

		/*
			// Update properties
			{
				Name:        "updateProperties",
				Usage:       "Update properties of an IP",
				Description: `ovh ip updateProperties IPBLOCK --desc "description"` + NLTAB + `Example: ovh ip updateProperties 37.187.0.144/32 --desc "IP routed to lunar base server"`,
				Flags: []cli.Flag{
					cli.StringFlag{"desc", "", "Update description", ""},
				},
				Action: func(c *cli.Context) {
					dieIfArgsMiss(len(c.Args()), 1)
					fDesc := c.String("desc")
					// check if there is something to update
					if len(fDesc) == 0 {
						dieDone()
					}
					err := ipr.UpdateProperties(c.Args().First(), fDesc)
					if err != nil {
						dieError(err)
					}
					dieDone()
				},
			},
			// Reverse
			{
				Name:        "reverse",
				Description: "commnands to interact with IP reverse",
				Subcommands: []cli.Command{
					{
						Name:        "get",
						Usage:       "Return the reverse of IP",
						Description: "ovh ip reverse XXX.XXX.XXX.XXX",
						Flags: []cli.Flag{
							cli.BoolFlag{Name: "json", Usage: "output as JSON"},
						},
						Action: func(c *cli.Context) {
							dieIfArgsMiss(len(c.Args()), 1)
							RIP, err := ipr.GetReverse(c.Args().First())
							if err != nil {
								dieError(err)
							}
							if c.Bool("json") {
								t, err := json.Marshal(RIP)
								if err != nil {
									dieError(err)
								}
								fmt.Println(string(t))
							} else {
								fmt.Println(RIP.String())
							}
							dieOk()
						},
					},
					{
						Name:        "set",
						Usage:       "Set the reverse of IP",
						Description: "ovh ip reverse set IP REVERSE",
						Action: func(c *cli.Context) {
							dieIfArgsMiss(len(c.Args()), 2)
							err := ipr.SetReverse(c.Args().First(), c.Args()[1])
							if err != nil {
								dieError(err)
							}
							dieOk()
						},
					},
				},
			},*/
	}
	return

}
