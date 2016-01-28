package main

import (
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
		}, /*{
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
				stats, err := ipr.SpamGetIPStats(ip.IPBlock{c.Args().First(), ""}, c.Args().Get(1), time.Unix(int64(from), 0), time.Unix(int64(to), 0))
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
				handleErrFromOvh(ipr.SpamUnblockSpamIP(ip.IPBlock{c.Args().First(), ""}, c.Args().Get(1)))
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
		},*/
	}
	return
}
