package main

import (
	"encoding/json"
	"fmt"

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
				print(formatOutput(properties, c.Bool("json")))
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
				if c.Bool("json") {
					r, _ := json.Marshal(tasks)
					dieOk(string(r))
				}
				for _, task := range tasks {
					fmt.Println(task)
				}
				dieOk()
			},
		},

		/*
			        	{
							Name:        "getTaskProperties",
							Usage:       "Return properties of a server task",
							Description: "ovh server getTaskProperties SERVER TASKID" + NLTAB + "Example: ovh server getTaskProperties ns309865.ovh.net 456",
							Action: func(c *cli.Context) {
								dieIfArgsMiss(len(c.Args()), 2)
								taskId, err := strconv.ParseUint(c.Args().Get(1), 10, 64)
								if err != nil {
									dieError(err)
								}
								task, err := sr.GetTaskProperties(c.Args().First(), taskId)
								dieOnError(err)
								fmt.Printf("Task ID: %d%s", task.Id, NL)
								fmt.Printf("Function: %s%s", task.Function, NL)
								fmt.Printf("Status: %s%s", task.Status, NL)
								fmt.Printf("Comment: %s%s", task.Comment, NL)
								fmt.Printf("Last Upadte: %s%s", task.LastUpdate, NL)
								fmt.Printf("Start Date: %s%s", task.StartDate, NL)
								fmt.Printf("Done Date: %s%s", task.DoneDate, NL)
								dieOk()
							},
						},
						{
							Name:        "cancelTask",
							Usage:       "Cancel a server task",
							Description: "ovh server cancelTask SERVER TASKID" + NLTAB + "Example: ovh server cancelTask ks323462.kimsufi.com 4319579",
							Action: func(c *cli.Context) {
								taskId, err := strconv.ParseUint(c.Args().Get(1), 10, 64)
								if err != nil {
									dieError(err)
								}
								err = sr.CancelTask(c.Args().Get(0), taskId)
								dieOnError(err)
								dieOk()
							},
						},
						{
							Name:        "reboot",
							Usage:       "Create a new reboot task",
							Description: "ovh server reboot SERVER" + NLTAB + "Example: ovh server reboot ks323462.kimsufi.com",
							Action: func(c *cli.Context) {
								dieIfArgsMiss(len(c.Args()), 1)
								task, err := sr.Reboot(c.Args().First())
								dieOnError(err)
								fmt.Printf("Task ID: %d%s", task.Id, NL)
								fmt.Printf("Function: %s%s", task.Function, NL)
								fmt.Printf("Status: %s%s", task.Status, NL)
								fmt.Printf("Comment: %s%s", task.Comment, NL)
								fmt.Printf("Last Upadte: %s%s", task.LastUpdate, NL)
								fmt.Printf("Start Date: %s%s", task.StartDate, NL)
								fmt.Printf("Done Date: %s%s", task.DoneDate, NL)
								dieOk()
							},
						},*/
	}

	/*
		switch cmd.Action {
		// List
		case "list":
			servers, err := serverR.List()
			if err != nil {
				return err
			}
			for _, s := range servers {
				resp = fmt.Sprintf("%s%s\r\n", resp, s)
			}
			if len(resp) > 2 {
				resp = resp[0 : len(resp)-2]
			}
			dieOk(resp)
			break
		// get server properties
		// ./ovh server getTasks SERVER_NAME FUNCTION STATUS
		case "properties":
			if len(cmd.Args) != 3 {
				return errors.New("\"server properties\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}

			properties, err := serverR.GetProperties(strings.ToLower(cmd.Args[2]))
			if err != nil {
				return err
			}
			fmt.Printf("ID: %d%s", properties.Id, NL)
			fmt.Printf("Name: %s%s", properties.Name, NL)
			fmt.Printf("Ip: %s%s", properties.Ip, NL)
			fmt.Printf("Datacenter: %s%s", properties.Datacenter, NL)
			fmt.Printf("ProfessionalUse: %t%s", properties.ProfessionalUse, NL)
			fmt.Printf("CommercialRange: %s%s", properties.CommercialRange, NL)
			fmt.Printf("Os: %s%s", properties.Os, NL)
			fmt.Printf("State: %s%s", properties.State, NL)
			fmt.Printf("Reverse: %s%s", properties.Reverse, NL)
			fmt.Printf("Monitored: %t%s", properties.Monitored, NL)
			fmt.Printf("Rack: %s%s", properties.Rack, NL)
			fmt.Printf("RootDevice: %s%s", properties.RootDevice, NL)
			fmt.Printf("LinkSpeed: %d%s", properties.LinkSpeed, NL)
			fmt.Printf("Bootid: %d%s", properties.BootId, NL)
			dieOk("")
			break

		// Get available netboots ID for this server
		case "availableNetboots":
			if len(cmd.Args) < 3 {
				return errors.New("\"server availableNetboots\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}
			var netbootIds []int
			if len(cmd.Args) == 3 {
				netbootIds, err = serverR.GetNetboots(strings.ToLower(cmd.Args[2]))
			} else {
				netbootIds, err = serverR.GetNetboots(strings.ToLower(cmd.Args[2]), cmd.Args[3])
			}
			if err != nil {
				return
			}
			for _, id := range netbootIds {
				fmt.Println(id)
			}
			break

		// Get server tasks
		case "getTasks":
			function := "all"
			status := "all"

			if len(cmd.Args) < 3 {
				return errors.New("\"server getTasks\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}

			if len(cmd.Args) > 5 {
				return errors.New("\"server getTasks\" too many arguments - see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}

			// serverName
			serverName := strings.ToLower(cmd.Args[2])

			// function
			if len(cmd.Args) > 3 {
				function = cmd.Args[3]
				if len(cmd.Args) > 4 {
					status = cmd.Args[4]
				}
			}

			tasks, err := serverR.GetTasks(serverName, function, status)
			if err != nil {
				return err
			}
			resp := ""
			for _, task := range tasks {
				resp = fmt.Sprintf("%s%d\r\n", resp, task)
			}
			if len(resp) > 2 {
				resp = resp[0 : len(resp)-2]
			}
			dieOk(resp)
			break

		// Get task properties
		// ./ovh server getTaskProperties SERVER_NAME TASK_ID
		case "getTaskProperties":
			if len(cmd.Args) != 4 {
				return errors.New("\"server getTaskProperties\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}

			serverName := strings.ToLower(cmd.Args[2])
			taskId, err := strconv.ParseUint(cmd.Args[3], 10, 64)
			if err != nil {
				return err
			}
			task, err := serverR.GetTaskProperties(serverName, taskId)
			if err != nil {
				return err
			}
			fmt.Printf("Task ID: %d%s", task.Id, NL)
			fmt.Printf("Function: %s%s", task.Function, NL)
			fmt.Printf("Status: %s%s", task.Status, NL)
			fmt.Printf("Comment: %s%s", task.Comment, NL)
			fmt.Printf("Last Upadte: %s%s", task.LastUpdate, NL)
			fmt.Printf("Start Date: %s%s", task.StartDate, NL)
			fmt.Printf("Done Date: %s%s", task.DoneDate, NL)
			dieOk("")
			break

		// Cancel task (if possible)
		// ./ovh server cancelTask SERVER_NAME TASK_ID
		case "cancelTask":
			if len(cmd.Args) != 4 {
				return errors.New("\"server cancel Task\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}
			serverName := strings.ToLower(cmd.Args[2])
			taskId, err := strconv.ParseUint(cmd.Args[3], 10, 64)
			if err != nil {
				return err
			}
			err = serverR.CancelTask(serverName, taskId)
			if err != nil {
				return err
			}
			dieOk("Success task " + cmd.Args[3] + "cancelled")
			break

		// Reboot
		case "reboot":
			if len(cmd.Args) != 3 {
				return errors.New("\"server reboot\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
			}
			task, err := serverR.Reboot(strings.ToLower(cmd.Args[2]))
			if err != nil {
				return err
			}
			fmt.Printf("Task ID: %d%s", task.Id, NL)
			fmt.Printf("Function: %s%s", task.Function, NL)
			fmt.Printf("Status: %s%s", task.Status, NL)
			fmt.Printf("Comment: %s%s", task.Comment, NL)
			fmt.Printf("Last Upadte: %s%s", task.LastUpdate, NL)
			fmt.Printf("Start Date: %s%s", task.StartDate, NL)
			fmt.Printf("Done Date: %s%s", task.DoneDate, NL)
			dieOk("")
			break

		default:
			return errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
		}*/
	return
}
