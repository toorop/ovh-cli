package main

import (
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/server"
	"strconv"
	"strings"
)

func serverHandler(cmd *Cmd) (err error) {
	// New govh client
	client := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)
	// New server ressource
	serverR, err := server.New(client)
	if err != nil {
		return
	}
	// response (string)
	var resp string

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
	}
	return
}
