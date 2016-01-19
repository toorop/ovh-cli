package main

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"

	"github.com/toorop/govh"
	"github.com/toorop/govh/cloud"
)

// getFwCmds return commands for firewall subsection
func getCloudCmds(client *govh.OVHClient) (cloudCmds []cli.Command) {
	cloud, err := cloud.New(client)
	if err != nil {
		return
	}
	cloudCmds = []cli.Command{
		{
			Name:        "getPassports",
			Usage:       "Return a list of cloud passports",
			Description: "Example: ovh cloud getPassport",
			Action: func(c *cli.Context) {
				passports, err := cloud.GetPassports()
				handleErrFromOvh(err)
				for _, passport := range passports {
					fmt.Println(passport)
				}
				dieOk()
			},
		}, {
			Name:        "getPrices",
			Usage:       "Return a list of cloud prices",
			Description: "Example: ovh cloud getPrices",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "if set output as JSON"},
			},
			Action: func(c *cli.Context) {
				prices, err := cloud.GetPrices()
				handleErrFromOvh(err)
				//	fmt.Println(prices.ProjectCreation)
				if c.Bool("json") {
					fmt.Println(prices.JSON())
				} else {
					fmt.Println(prices.String())
				}
				dieOk()
			},
		}, {

			Name:        "getProjectsId",
			Usage:       "Return a list of projects ID",
			Description: "Example: ovh cloud getProjectsId",
			Action: func(c *cli.Context) {
				ids, err := cloud.GetProjectsId()
				handleErrFromOvh(err)
				for _, id := range ids {
					fmt.Println(id)
				}
				dieOk()
			},
		}, {
			Name:        "getProject",
			Usage:       "Return project info",
			Description: "Example: ovh cloud getProject PROJECT_ID",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				project, err := cloud.GetProject(c.Args().First())
				handleErrFromOvh(err)
				fmt.Printf("Project Id: %s%s", project.Id, NL)
				fmt.Printf("Status: %s%s", project.Status, NL)
				fmt.Printf("Creation date: %s%s", project.CreationDate.Format(time.RFC3339), NL)
				fmt.Printf("Description: %s%s", project.Description, NL)
				dieOk()
			},
		},
	}
	return
}
