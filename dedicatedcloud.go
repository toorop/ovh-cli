package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/dedicatedcloud"
)

// getDedicatedCloudCmds return commands for Dedicated Cloud subsection
func getDedicatedCloudCmds(client *govh.OVHClient) (dedicatedCloudCmds []cli.Command) {
	sr, err := dedicatedcloud.New(client)
	if err != nil {
		return
	}

	dedicatedCloudCmds = []cli.Command{
		{
			Name:        "list",
			Usage:       "Return a list of Dedicated Cloud",
			Description: "ovh dedicatedcloud list [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud list",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dedicatedClouds, err := sr.List()
				dieOnError(err)
				for _, dedicatedCloud := range dedicatedClouds {
					fmt.Println(dedicatedCloud)
				}
				dieOk()
			},
		},
		{
			Name:        "properties",
			Usage:       "Return properties of a Dedicated Cloud",
			Description: "ovh dedicatedcloud properties DEDICATEDCLOUD [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud properties pcc-123-123-123-123",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				properties, err := sr.GetProperties(c.Args().First())
				dieOnError(err)
				println(formatOutput(properties, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "users",
			Usage:       "Return a list of users for a Dedicated Cloud",
			Description: "ovh dedicatedcloud users DEDICATEDCLOUD [--name NAME] [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud users pcc-123-123-123-123 --name admin",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "name", Value: "", Usage: "(optional) : filter by user name."},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				users, err := sr.GetUsers(c.Args().First(), c.String("name"))
				dieOnError(err)
				println(formatOutput(users, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "user",
			Usage:       "Return properties of a Dedicated Cloud user",
			Description: "ovh dedicatedcloud user DEDICATEDCLOUD USERID [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud user pcc-123-123-123-123 456",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				userID, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
				if err != nil {
					dieError(err)
				}
				user, err := sr.GetUser(c.Args().First(), int(userID))
				dieOnError(err)
				println(formatOutput(user, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "datacenters",
			Usage:       "Return a list of datacenters for a Dedicated Cloud",
			Description: "ovh dedicatedcloud datacenters DEDICATEDCLOUD [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud datacenters pcc-123-123-123-123",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				datacenters, err := sr.GetDatacenters(c.Args().First())
				dieOnError(err)
				println(formatOutput(datacenters, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "datacenter",
			Usage:       "Return properties of a Dedicated Cloud datacenter",
			Description: "ovh dedicatedcloud datacenter DEDICATEDCLOUD DATACENTERID [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud datacenter pcc-123-123-123-123 456",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				datacenterID, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
				if err != nil {
					dieError(err)
				}
				datacenter, err := sr.GetDatacenter(c.Args().First(), int(datacenterID))
				dieOnError(err)
				println(formatOutput(datacenter, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "tasks",
			Usage:       "Return a list of tasks for a Dedicated Cloud",
			Description: "ovh dedicatedcloud tasks DEDICATEDCLOUD [--state STATE] [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud tasks pcc-123-123-123-123 --state done",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "state", Value: "", Usage: "(optional) : filter by state. See [OVH doc](https://api.ovh.com/console/#/dedicatedCloud/%7BserviceName%%7D/task#GET) for availables states."},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				tasks, err := sr.GetTasks(c.Args().First(), c.String("state"))
				dieOnError(err)
				println(formatOutput(tasks, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "task",
			Usage:       "Return properties of a Dedicated Cloud task",
			Description: "ovh dedicatedcloud task DEDICATEDCLOUD TASKID [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud task pcc-123-123-123-123 456",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				taskID, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
				if err != nil {
					dieError(err)
				}
				task, err := sr.GetTask(c.Args().First(), int(taskID))
				dieOnError(err)
				println(formatOutput(task, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "allowednetworks",
			Usage:       "Return a list of AllowedNetwork for a Dedicated Cloud",
			Description: "ovh dedicatedcloud tasks DEDICATEDCLOUD [--state STATE] [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud allowednetworks pcc-123-123-123-123",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				tasks, err := sr.GetAllowedNetworks(c.Args().First())
				dieOnError(err)
				println(formatOutput(tasks, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "allowednetwork",
			Usage:       "Return properties of a Dedicated Cloud allowednetworks",
			Description: "ovh dedicatedcloud allowednetwork DEDICATEDCLOUD ALLOWEDNETWORKID [--json]" + NLTAB + TAB + "Example: ovh dedicatedcloud allowednetwork pcc-123-123-123-123 456",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				allowedNetworkID, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
				if err != nil {
					dieError(err)
				}
				task, err := sr.GetAllowedNetwork(c.Args().First(), int(allowedNetworkID))
				dieOnError(err)
				println(formatOutput(task, c.Bool("json")))
				dieOk()
			},
		},
	}
	return
}
