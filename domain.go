package main

import (
	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/domain"
)

// getIpCmds return commands for Ip section
func getDomainCmds(OVHClient *govh.OVHClient) (cmds []cli.Command) {
	domClient, err := domain.New(OVHClient)
	if err != nil {
		return
	}

	// Ip commands
	cmds = []cli.Command{
		// IPBLock
		{
			Name:        "list",
			Description: "list domain (all or filter by whois owner)",
			Usage:       "ovh domain list [--owner WHOISOWNER] [--json]" + NLTAB + "Example: ovh domaink list --owner XXXX --json",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "owner", Value: "", Usage: "Filter by whois owner"},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				domains, err := domClient.List(c.String("owner"))
				dieOnError(err)
				if c.Bool("json") {
					buf, err := json.Marshal(domains)
					dieOnError(err)
					fmt.Println(string(buf))
				} else {
					for _, domain := range domains {
						fmt.Println(domain)
					}
				}
			},
		}, {
			Name:        "zone",
			Description: "subcomands for DNS zones",
			Subcommands: []cli.Command{
				// List Ip Blocks
				{
					Name:        "getrecordid",
					Description: "returns record IDs for a zone",
					Usage:       "ovh domain zone getrecordid ZONE [--field FIELD] [--sub SUBDOMAIN] [--json]" + NLTAB + "Example: ovh domain zone getrecordid ovh.com --field A --json",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "field", Value: "", Usage: "Filter by DNS field type (A, MX, TXT,...)"},
						cli.StringFlag{Name: "sub", Value: "", Usage: "Filter by subdomain"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						IDs, err := domClient.GetRecordIDs(c.Args().First(), domain.GetRecordIDsOptions{
							FieldType: c.String("field"),
							SubDomain: c.String("sub"),
						})
						dieOnError(err)
						if c.Bool("json") {
							buf, err := json.Marshal(IDs)
							dieOnError(err)
							fmt.Println(string(buf))
						} else {
							for _, ID := range IDs {
								fmt.Println(ID)
							}
						}
					},
				},
			},
		},
	}
	return
}
