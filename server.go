package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/dedicated/server"
)

// getFwCmds return commands for firewall subsection
func getServerCmds(client *govh.OVHClient) (serverCmds []cli.Command) {
	sr, err := server.New(client)
	if err != nil {
		return
	}

	serverCmds = []cli.Command{
		{
			Name:        "list",
			Usage:       "Return a list of server ",
			Description: "ovh server list" + NLTAB + "Example: ovh server list",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				servers, err := sr.List()
				dieOnError(err)
				for _, server := range servers {
					fmt.Println(server)
				}
				dieOk()
			},
		},
		{
			Name:        "properties",
			Usage:       "Return properties of a server ",
			Description: "ovh server properties SERVER [--json]" + NLTAB + "Example: ovh server properties ks323462.kimsufi.com",
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
			Name:        "tasks",
			Usage:       "Return a list of tasks for a server ",
			Description: "ovh server tasks SERVER [--function FUNCTION] [--status STATUS] [--json]" + NLTAB + "Example: ovh server tasks ns309865.ovh.net --function hardReboot --status done --json",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "function", Value: "", Usage: "(optional) - filter by function. See https://api.ovh.com/console/#/dedicated/server/%7BserviceName%%7D/task#GET for availables functions.)"},
				cli.StringFlag{Name: "status", Value: "", Usage: "(optional) : filter by status. See [OVH doc](https://api.ovh.com/console/#/dedicated/server/%7BserviceName%%7D/task#GET) for availables status."},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				tasks, err := sr.GetTasks(c.Args().First(), c.String("function"), c.String("status"))
				dieOnError(err)
				println(formatOutput(tasks, c.Bool("json")))
				dieOk()
			},
		},

		{
			Name:        "task",
			Usage:       "Return properties of a server task",
			Description: "ovh server task SERVER TASKID" + NLTAB + "Example: ovh server task ns309865.ovh.net 456",
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
			Name:        "canceltask",
			Usage:       "Cancel a server task",
			Description: "ovh server canceltask SERVER TASKID" + NLTAB + "Example: ovh server canceltask ks323462.kimsufi.com 4319579",
			Action: func(c *cli.Context) {
				taskID, err := strconv.ParseUint(c.Args().Get(1), 10, 64)
				if err != nil {
					dieError(err)
				}
				err = sr.CancelTask(c.Args().Get(0), int(taskID))
				dieOnError(err)
				dieOk()
			},
		},

		{
			Name:        "reboot",
			Usage:       "Create a new reboot task",
			Description: "ovh server reboot SERVER [--json]" + NLTAB + "Example: ovh server reboot ks323462.kimsufi.com --json",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				task, err := sr.Reboot(c.Args().First())
				dieOnError(err)
				println(formatOutput(task, c.Bool("json")))
				dieOk()
			},
		},
	}
	return
}
