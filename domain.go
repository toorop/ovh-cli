package main

import (
	"io/ioutil"
	"os"
	"strconv"

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

	// /domain commands
	cmds = []cli.Command{
		// list domains
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
				println(formatOutput(domains, c.Bool("json")))
			},
			// Zone
		}, {
			Name:        "zone",
			Description: "subcomands for DNS zones",
			Subcommands: []cli.Command{
				// Create record
				{
					Name:        "newrecord",
					Description: "creates a new record",
					Usage:       "ovh domain zone newrecord ZONE --field FIELD --target TARGET [--ttl TTL] [--sub SUBDOMAIN] [--json]" + NLTAB + "Example: ovh domain zone newrecord ovh.com --field A --target 8.8.8.8 --ttl 300 --sub ovhcli --json",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "field", Value: "", Usage: "DNS field type (A, MX, TXT,...)"},
						cli.StringFlag{Name: "target", Value: "", Usage: "DNS target (eg 127.0.0.1)"},
						cli.StringFlag{Name: "ttl", Value: "0", Usage: "TTL"},
						cli.StringFlag{Name: "sub", Value: "", Usage: "sub domain"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						ttl, err := strconv.ParseInt(c.String("ttl"), 10, 32)
						dieOnError(err)

						record, err := domClient.NewRecord(domain.ZoneRecord{
							Zone:      c.Args().First(),
							Target:    c.String("target"),
							TTL:       int(ttl),
							FieldType: c.String("field"),
							SubDomain: c.String("sub"),
						})
						dieOnError(err)
						println(formatOutput(record, c.Bool("json")))
						dieOk()
					},
				},
				// List records ID
				{
					Name:        "getrecordsid",
					Description: "returns record IDs for a zone",
					Usage:       "ovh domain zone getrecordsid ZONE [--field FIELD] [--sub SUBDOMAIN] [--json]" + NLTAB + "Example: ovh domain zone getrecordsid ovh.com --field A --json",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "field", Value: "", Usage: "Filter by DNS field type (A, MX, TXT,...)"},
						cli.StringFlag{Name: "sub", Value: "", Usage: "Filter by subdomain"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						IDs, err := domClient.GetRecordIDs(c.Args().First(), domain.GetRecordsOptions{
							FieldType: c.String("field"),
							SubDomain: c.String("sub"),
						})
						dieOnError(err)
						println(formatOutput(IDs, c.Bool("json")))
					},
				},

				// Get record
				{
					Name:        "getrecords",
					Description: "returns records for a zone",
					Usage:       "ovh domain zone getrecords ZONE [--field FIELD] [--sub SUBDOMAIN] [--json]" + NLTAB + "Example: ovh domain zone getrecords ovh.com --field A --json",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "field", Value: "", Usage: "Filter by DNS field type (A, MX, TXT,...)"},
						cli.StringFlag{Name: "sub", Value: "", Usage: "Filter by subdomain"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						records, err := domClient.GetRecords(c.Args().First(), domain.GetRecordsOptions{
							FieldType: c.String("field"),
							SubDomain: c.String("sub"),
						})
						dieOnError(err)
						println(formatOutput(records, c.Bool("json")))
					},
				},

				// Delete record
				{
					Name:        "delrecord",
					Description: "delete record",
					Usage:       "ovh domain zone delrecord ZONE RECORD_ID" + NLTAB + "Example: ovh domain zone delrecord ovh.com 123456",
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 2)
						ID, err := strconv.ParseInt(c.Args()[1], 10, 32)
						dieOnError(err)
						dieOnError(domClient.DeleteRecord(c.Args()[0], int(ID)))
						dieOk()
					},
				},
				// get (import) zone file
				{
					Name:        "import",
					Description: "import zone as Bind zonefile format",
					Usage:       "ovh domain zone import ZONE" + NLTAB + "Example: ovh domain zone import ovh.com",
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						zoneFile, err := domClient.GetZoneFile(c.Args().First())
						dieOnError(err)
						dieOk(zoneFile)
					},
				},
				// put (export) zone file
				{
					Name:        "export",
					Description: "export zone as Bind format to OVH",
					Usage:       "ovh domain zone export ZONE [--zone PATH_TO_FILE]" + NLTAB + "Example: ovh domain zone  ovh.com --zone /tmp/domain.com.zone",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "zone", Value: "", Usage: "path to zone file"},
						cli.BoolFlag{Name: "json", Usage: "output as JSON"},
					},
					Action: func(c *cli.Context) {
						var zoneFile []byte
						var err error
						dieIfArgsMiss(len(c.Args()), 1)
						zoneFilePath := c.String("zone")
						if zoneFilePath != "" {
							zoneFile, err = ioutil.ReadFile(zoneFilePath)
						} else {
							zoneFile, err = ioutil.ReadAll(os.Stdin)
						}
						dieOnError(err)
						task, err := domClient.PutZoneFile(c.Args().First(), string(zoneFile))
						dieOnError(err)
						dieOk(formatOutput(task, c.Bool("json")))
					},
				},
				// Refresh DNSes
				{
					Name:        "reload",
					Description: "apply zone modifications on DNS servers",
					Usage:       "ovh domain zone reload ZONE" + NLTAB + "Example: ovh domain zone reload ovh.com",
					Action: func(c *cli.Context) {
						dieIfArgsMiss(len(c.Args()), 1)
						dieOnError(domClient.RefreshZone(c.Args()[0]))
						dieOk()
					},
				},
			},
		},
	}
	return
}
