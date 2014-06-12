package main

import (
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	"github.com/codegangsta/cli"
	"strings"
)

// getIpCmds return commands for Ip section
func getIpCmds(client *govh.OvhClient) (ipCmds []cli.Command) {
	ipCmds = []cli.Command{
		// list
		{
			Name:        "list",
			Usage:       "List your IP blocks.",
			Description: "ovh ip list [flag...]" + NLTAB + "Example: ovh ip list --type vps",
			Flags: []cli.Flag{
				cli.StringFlag{"desc", "", "Filter: by description (like)."},
				cli.StringFlag{"ip", "", "Filter: by IP (contains or equals)."},
				cli.StringFlag{"routedTo", "", "Filter: by routing."},
				cli.StringFlag{"type", "all", "Filter: by IP block type: all|cdn|dedicated|failover|hosted_ssl|housing|loadBalancing|mail|pcc|pci|private|vps|vpn|vrack|xdsl"},
			},
			Action: func(c *cli.Context) {
				fDesc := strings.ToLower(c.String("desc"))
				fIp := strings.ToLower(c.String("ip"))
				fRoutedTo := strings.ToLower(c.String("routedto"))
				fType := strings.ToLower(c.String("type"))
				if fType == "all" {
					fType = ""
				}

				if ipr, err = ip.New(client); err != nil {
					return
				}
				ips, err := ipr.List(fDesc, fIp, fRoutedTo, fType)
				handleErrFromOvh(err)
				for _, i := range ips {
					fmt.Println(i.IP, i.Type)
				}
				dieOk()
			},
		},
		// getProperties
		{
			Name:        "getProperties",
			Usage:       "Get properties of an IP.",
			Description: "ovh ip getProperties IPBLOCK" + NLTAB + "Example: ovh ip getProperties 91.121.228.135/32",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				if ipr, err = ip.New(client); err != nil {
					return
				}
				properties, err := ipr.GetIpProperties(c.Args().First())
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
				cli.StringFlag{"desc", "", "Update description"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				fDesc := c.String("desc")
				// check if there is something to update
				if len(fDesc) == 0 {
					dieDone()
				}
				if ipr, err = ip.New(client); err != nil {
					return
				}
				err := ipr.UpdateProperties(c.Args().First(), fDesc)
				if err != nil {
					dieError(err)
				}
				dieDone()
			},
		},
	}
	return

}
