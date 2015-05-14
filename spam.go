package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/ip"
	"time"
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
			Usage:       "List IP which send (or have sent) spam.",
			Description: "ovh spam listIp IPBLOCK [--state ]" + NLTAB + "Example: ovh spam listIp 91.121.228.135/32 --state unblocked",
			Flags: []cli.Flag{
				cli.StringFlag{"state", "", "The state of the IP (blockedForSpam|unblocked|unblocking).", ""},
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
		}, {
			Name:        "getProperties",
			Usage:       "Get properties of a spamming IP.",
			Description: "ovh spam getProperties IPBLOCK IP" + NLTAB + "Example: ovh spam listIp 91.121.228.135/32 91.121.228.135",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				p, err := ipr.SpamGetIp(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1))
				handleErrFromOvh(err)
				dieOk(fmt.Sprintf("Blocked since (duration sec):%d%sLast time: %s%sIP: %s%sState: %s", p.Time, NL, p.Date, NL, p.IpSpamming, NL, p.State))
			},
		}, {
			Name:        "getStats",
			Usage:       "Get spam stats about a spamming IP.",
			Description: "ovh spam getStats IPBLOCK IP --from TIMESTAMP_START --to TIMESTAMP_STOP" + NLTAB + "Example: ovh spam getStats 178.33.223.32/28 178.33.223.42 --from 1385251200 --to 1387882630",
			Flags: []cli.Flag{
				cli.StringFlag{"from", "", "Unix timestamp representing the begining of the peiod (required).", ""},
				cli.StringFlag{"to", "", "Unix timestamp representing the end of the peiod (required).", ""},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				if !c.IsSet("from") || !c.IsSet("to") {
					dieBadArgs()
				}
				from := c.Int("from")
				to := c.Int("to")
				if from >= to {
					dieBadArgs()
				}
				stats, err := ipr.SpamGetIpStats(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1), time.Unix(int64(from), 0), time.Unix(int64(to), 0))
				handleErrFromOvh(err)
				if stats == nil {
					dieOk("No spam stats for this period")
				}
				fmt.Printf("Blocked for the last time: %s%s", time.Unix(stats.Timestamp, 0).Format(time.RFC822Z), NL)
				fmt.Printf("Number of emails sent: %d%s", stats.Total, NL)
				fmt.Printf("Number of spams sent: %d%s", stats.NumberOfSpams, NL)
				fmt.Printf("Average score: %d%s%s", stats.AverageSpamScore, NL, NL)
				if len(stats.DetectedSpams) > 0 {
					fmt.Println("Detected Spams : ", NL)
				}
				for _, ds := range stats.DetectedSpams {
					fmt.Println("")
					fmt.Printf("\tDate: %s%s", time.Unix(ds.Date, 0).Format(time.RFC822Z), NL)
					fmt.Printf("\tMessage ID: %s%s", ds.MessageId, NL)
					fmt.Printf("\tDestination IP: %s%s", ds.DestinationIp, NL)
					fmt.Printf("\tScore: %d%s", ds.Spamscore, NL)
				}
				dieOk()
			},
		},
		{
			Name:        "unblock",
			Usage:       "Unblock a locked IP.",
			Description: "ovh spam unblock IPBLOCK IP" + NLTAB + "Example: ovh spam unblock 91.121.228.135/32 91.121.228.135",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				handleErrFromOvh(ipr.SpamUnblockSpamIp(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1)))
				dieDone()
			},
		},
		{
			Name:        "getBlocked",
			Usage:       "Retuns IPs which are currently blocked.",
			Description: "ovh spam getBlocked" + NLTAB + "Example: ovh spam getBlocked",
			Action: func(c *cli.Context) {
				//dieIfArgsMiss(len(c.Args()), 2)
				ips, err := ipr.GetBlockedForSpam()
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
