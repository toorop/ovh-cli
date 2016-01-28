package main

import (
	"time"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/me"
)

// getMeCmds return commands for /me section
func getMeCmds(OVHClient *govh.OVHClient) (cmds []cli.Command) {
	meClient, err := me.New(OVHClient)
	if err != nil {
		return
	}

	// Ip commands
	cmds = []cli.Command{
		{
			// SUB bill
			Name:        "bill",
			Description: "subcomands for bill",
			Subcommands: []cli.Command{
				{
					// CMD list - list bill ID
					Name:        "list",
					Description: "return bill IDs from dateFrom to dateTo",
					Usage:       "ovh me bill list [--from TIMESTAMP] [--to TIMESTAMP] [--json]" + NLTAB + "Example: ovh me bill list --from 1420066800 --to 1451602800",
					Flags: []cli.Flag{
						cli.IntFlag{Name: "from", Value: 0, Usage: "Date from"},
						cli.IntFlag{Name: "to", Value: 0, Usage: "Date to"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						var dateFrom, dateTo time.Time
						dateFrom = time.Unix(int64(c.Int("from")), 0)
						if c.Int("to") == 0 {
							dateTo = time.Now()
						} else {
							dateTo = time.Unix(int64(c.Int("to")), 0)
						}
						IDs, err := meClient.GetBillIDs(dateFrom, dateTo)
						dieOnError(err)
						println(formatOutput(IDs, c.Bool("json")))
						dieOk()
					},
				}, {
					// CMD getbyid - returns bill by its ID
					Name:        "getbyid",
					Description: "returns bill from its ID",
					Usage:       "ovh me bill getbyid ID [--json]" + NLTAB + "Example: ovh me bill getbyid 123456789 --json",
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						bill, err := meClient.GetBillByID(c.Args().First())
						dieOnError(err)
						println(formatOutput(bill, c.Bool("json")))
						dieOk()
					},
				}, {
					// CMD dowload - download bills as PDF
					Name:        "download",
					Description: "download bills from dateFrom to dateTo and save them to path",
					Usage:       "ovh me bill download --path SAVEPATH [--from TIMESTAMP] [--to TIMESTAMP] [--json]" + NLTAB + "Example: ovh me bill list --path /tmp --from 1420066800 --to 1451602800",
					Flags: []cli.Flag{
						cli.IntFlag{Name: "from", Value: 0, Usage: "Date from"},
						cli.IntFlag{Name: "to", Value: 0, Usage: "Date to"},
					},

					Action: func(c *cli.Context) {
						// TODO check path

						var dateFrom, dateTo time.Time
						dateFrom = time.Unix(int64(c.Int("from")), 0)
						if c.Int("to") == 0 {
							dateTo = time.Now()
						} else {
							dateTo = time.Unix(int64(c.Int("to")), 0)
						}
						_, err := meClient.GetBillIDs(dateFrom, dateTo)
						dieOnError(err)
						// TODO DL && save files

						dieOk()
					},
				},
			},
		},
	}
	return
}
