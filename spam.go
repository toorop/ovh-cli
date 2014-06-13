package main

import (
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	"github.com/codegangsta/cli"
)

// getFwCmds return commands for firewall subsection
func getSpamCmds(client *govh.OvhClient) (spamCmds []cli.Command) {
	ipr, err := ip.New(client)
	if err != nil {
		return
	}

	spamCmds = []cli.Command{
		{
			Name:        "listIp",
			Usage:       "List IP which send (or have sent) spam",
			Description: "ovh spam listIp IPBLOCK [--state ]" + NLTAB + "Example: ovh spam listIp 91.121.228.135/32 --state unblocked",
			Flags: []cli.Flag{
				cli.StringFlag{"state", "", "The state of the IP (blockedForSpam|unblocked|unblocking)."},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				var state string
				if c.IsSet("state") {
					state = c.String("state")
					if !inSliceStr(state, []string{"blockedForSpam", "unblocked", "unblocking"}) {
						dieBadArgs()
					}
				}
				ips, err := ipr.SpamGetSpammingIps(ip.IpBlock{c.Args().First(), ""}, state)
				handleErrFromOvh(err)
				for _, ip := range ips {
					fmt.Println(ip)
				}
				dieOk()
			},
		},
	}
	return
}
