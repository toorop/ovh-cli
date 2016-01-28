package main

import (
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
		// IP
		// Get IP reverse
		{
			Name:        "reverse",
			Description: "Returns the reverse of the IP",
			Usage:       "ovh ip reverse [--json]" + NLTAB + "Example: ovh ip reverse --json",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				RIP, err := IPClient.GetReverse(c.Args().First())
				dieOnError(err)
				println(formatOutput(RIP, c.Bool("json")))
				dieOk()
			},
		}, {
			Name:        "setreverse",
			Description: "Update the reverse of the IP",
			Usage:       "ovh ip setreverse IP REVERSE [--json]" + NLTAB + "Example: ovh ip setreverse 8.8.8.8 www.ovh.com",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "reverse", Value: "", Usage: "new reverse for IP"},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				dieOnError(IPClient.SetReverse(c.Args().First(), c.Args()[1]))
			},
		},

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
						println(formatOutput(IPBlocks, c.Bool("json")))
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
						println(formatOutput(properties, c.Bool("json")))
						dieOk()
					},
				},
				// Update properties of a block
				// for know tou can only update description
				{
					Name:        "updateproperties",
					Usage:       "Update properties of an IP Block",
					Description: `ovh ip block updateproperties IPBLOCK --desc "description"` + NLTAB + `Example: ovh ip block updateproperties 37.187.0.144/32 --desc "Block routed to lunar base server"`,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "desc", Value: "", Usage: "Update block description"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						dieOnError(IPClient.UpdateBlockProperties(c.Args().First(), c.String("desc")))
						dieOk()
					},
				}, /*{
					Name:        "getreverses",
					Usage:       "Return the reverse of IP",
					Description: "ovh ip reverse XXX.XXX.XXX.XXX",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						RIP, err := IPClient.GetReverse(c.Args().First())
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
				},*/
			},
		}, // end of block subCommands

		/*
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
							RIP, err := IPClient.GetReverse(c.Args().First())
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
