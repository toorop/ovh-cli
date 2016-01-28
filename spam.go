package main

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/ip"
)

// getFwCmds return commands for firewall subsection
func getSpamCmds(client *govh.OVHClient) (cmds []cli.Command) {
	spamClient, err := ip.New(client)
	if err != nil {
		return
	}

	cmds = []cli.Command{
		{
			Name:        "list",
			Usage:       "List IP which send (or have sent) spam.",
			Description: "ovh spam list IPBLOCK [--state ] [--json]" + NLTAB + "Example: ovh spam list 91.121.228.135/32 --state unblocked",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "state", Value: "", Usage: "The state of the IP (blockedForSpam|unblocked|unblocking)."},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				ips, err := spamClient.SpamGetIPs(ip.IPBlock(c.Args().First()), c.String("state"))
				dieOnError(err)
				println(formatOutput(ips, c.Bool("json")))
				dieOk()
			},
		}, {
			Name:        "getproperties",
			Usage:       "Get properties of a spamming IP.",
			Description: "ovh spam getproperties IPBLOCK IP [--json]" + NLTAB + "Example: ovh spam getproperties 91.121.228.135/32 91.121.228.135",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				p, err := spamClient.SpamGetIP(ip.IPBlock(c.Args().First()), c.Args().Get(1))
				dieOnError(err)
				println(formatOutput(p, c.Bool("json")))
				dieOk()

			},
		}, {
			Name:        "stats",
			Usage:       "Get spam stats about a spamming IP.",
			Description: "ovh spam stats IPBLOCK IP --from TIMESTAMP_START --to TIMESTAMP_STOP" + NLTAB + "Example: ovh spam getStats 178.33.223.32/28 178.33.223.42 --from 1385251200 --to 1387882630",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "from", Value: "", Usage: "Unix timestamp representing the begining of the peiod (required)."},
				cli.StringFlag{Name: "to", Value: "", Usage: "Unix timestamp representing the end of the peiod (required)."},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				var from, to time.Time
				from = time.Unix(int64(c.Int("from")), 0)
				if c.Int("to") == 0 {
					to = time.Now()
				} else {
					to = time.Unix(int64(c.Int("to")), 0)
				}
				stats, err := spamClient.SpamGetIPStats(ip.IPBlock(c.Args().First()), c.Args().Get(1), from, to)
				dieOnError(err)
				dieOk(formatOutput(stats, c.Bool("json")))
			},
		},
		{
			Name:        "unblock",
			Usage:       "Unblock a blocked IP.",
			Description: "ovh spam unblock IPBLOCK IP" + NLTAB + "Example: ovh spam unblock 91.121.228.135/32 91.121.228.135",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				dieOnError(spamClient.SpamUnblockSpamIP(ip.IPBlock(c.Args().First()), c.Args().Get(1)))
				dieOk()
			},
		},

		{
			Name:        "getblocked",
			Usage:       "Retuns IPs which are currently blocked.",
			Description: "ovh spam getblocked" + NLTAB + "Example: ovh spam getblocked",
			Action: func(c *cli.Context) {
				ips, err := spamClient.GetBlockedForSpam()
				dieOnError(err)
				for _, ip := range ips {
					fmt.Println(ip)
				}
				dieOk()
			},
		},
	}
	return
}
