package main

import (
	//"encoding/json"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/sms"
	"strings"
	//"time"
)

func smsHandler(cmd *Cmd) (err error) {
	// New govh client
	client := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)
	// New sms ressource
	smsr, err := sms.New(client)

	switch cmd.Action {
	// List SMS servives
	// ./ovh sms listServices
	case "listServices":
		services, err := smsr.ListServices()
		if err != nil {
			dieError(err)
		}
		for _, s := range services {
			fmt.Println(s)
		}
		dieOk("")
	// Send a new SMS
	// ./ovh sms SERVICE_NAME new {JSON encoded SMS job}
	case "new":
		var job sms.NewJob
		err := json.Unmarshal([]byte(cmd.Args[3]), &job)
		if err != nil {
			dieError(err)
		}
		report, err := smsr.AddJob(cmd.Args[2], &job)
		if err != nil {
			dieError(err)
		}
		fmt.Printf("Done%s", NL)
		for _, id := range report.Ids {
			fmt.Printf("Job ID: %d%s", id, NL)
		}
		fmt.Printf("Credits removed: %d%s", report.TotalCreditsRemoved, NL)
		dieOk("")
	default:
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
	}
	return

}
