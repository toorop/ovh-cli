package main

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/ip"
)

// getIpCmds return commands for Ip section
func getIPCmds(client *govh.OVHClient) (ipCmds []cli.Command) {
	ipr, err := ip.New(client)
	if err != nil {
		return
	}

	// Ip commands
	ipCmds = []cli.Command{
		// IPBLock
		{
			Name:        "block",
			Description: "commands concerning IP block",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Description: "List your IP blocks.",
					Usage:       "ovh ip block list [flag...]" + NLTAB + "Example: ovh ip block list --type vps",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "desc", Value: "", Usage: "Filter: by description (like)."},
						cli.StringFlag{Name: "ip", Value: "", Usage: "Filter: by IP (contains or equals)."},
						cli.StringFlag{Name: "routedTo", Value: "", Usage: "Filter: by routing."},
						cli.StringFlag{Name: "type", Value: "all", Usage: "Filter: by IP block type: all|cdn|dedicated|failover|hosted_ssl|housing|loadBalancing|mail|pcc|pci|private|vps|vpn|vrack|xdsl"},
					},
					Action: func(c *cli.Context) {
						fDesc := strings.ToLower(c.String("desc"))
						fIp := strings.ToLower(c.String("ip"))
						fRoutedTo := strings.ToLower(c.String("routedto"))
						fType := strings.ToLower(c.String("type"))
						if fType == "all" {
							fType = ""
						}

						ips, err := ipr.List(fDesc, fIp, fRoutedTo, fType)
						handleErrFromOvh(err)
						for _, i := range ips {
							fmt.Println(i)
						}
						dieOk()
					},
				},
			},
		},

		/*
			// getProperties
			{
				Name:        "getProperties",
				Usage:       "Get properties of an IP.",
				Description: "ovh ip getProperties IPBLOCK" + NLTAB + "Example: ovh ip getProperties 91.121.228.135/32",
				Action: func(c *cli.Context) {
					dieIfArgsMiss(len(c.Args()), 1)
					properties, err := ipr.GetIPProperties(c.Args().First())
					handleErrFromOvh(err)
					dieOk(fmt.Sprintf("IP: %s%sType: %s%sDescription: %s%sRouted to: %s", properties.Ip, NL, properties.Type, NL, properties.Description, NL, properties.RoutedTo.ServiceName))
				},
			},

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
