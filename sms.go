package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/sms"
	"strings"
)

// getFwCmds return commands for firewall subsection
func getSmsCmds(client *govh.OvhClient) (smsCmds []cli.Command) {
	sr, err := sms.New(client)
	if err != nil {
		return
	}
	smsCmds = []cli.Command{
		{
			Name:        "listServices",
			Usage:       "Return a list of sms services ",
			Description: "ovh sms listServices" + NLTAB + "Example: ovh sms listServices",
			Action: func(c *cli.Context) {
				services, err := sr.ListServices()
				handleErrFromOvh(err)
				for _, service := range services {
					fmt.Println(service)
				}
				dieOk()
			},
		},
		{
			Name:        "send",
			Usage:       "Send an new SMS",
			Description: `ovh sms send SERVICE [--flags...]" + NLTAB + "Example: ovh sms send sms-st2-1 --sender +33979XXXX --receivers +336222XXXX +336221XXXX --message "Test from ovh-cli"`,
			Flags: []cli.Flag{
				cli.StringFlag{"sender", "", "The sender phone number in international format (+33XXXXXX for France for ex). Required.", ""},
				cli.StringSliceFlag{"receiver", &cli.StringSlice{}, "Receiver phone number. If you have multiple receivers add on --receiver flag by reciever. Requiered.", ""},
				cli.StringFlag{"message", "", "The message you want to send. Required", ""},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)

				// sender
				sender := c.String("sender")
				if sender == "" {
					dieBadArgs("Required flag --sender is missing")
				}
				// recievers
				receivers := c.StringSlice("receiver")
				if receivers == nil {
					dieBadArgs("Required flag --receivers is missing")
				}
				/*for _, rcv := range receivers {
					fmt.Println(rcv)
				}*/

				// message
				message := c.String("message")
				if message == "" {
					dieBadArgs("Required flag --message is missing")
				}

				// ValidityPeriod
				validityPeriod := 2880

				// Class
				class := "sim"

				// Create the job
				job := &sms.SendJob{
					Sender:         sender,
					Receivers:      receivers,
					Message:        message,
					ValidityPeriod: validityPeriod,
					Class:          class,
				}
				resp, err := sr.AddJob(c.Args().First(), job)
				handleErrFromOvh(err)
				for _, id := range resp.Ids {
					fmt.Println("Job ID:", id)
				}
				fmt.Println("Invalid receivers:", strings.Join(resp.InvalidReceivers, ", "))
				fmt.Println("Valid receivers:", strings.Join(resp.ValidReceivers, ", "))
				fmt.Println("Credits removed:", fmt.Sprintf("%d", resp.TotalCreditsRemoved))
				dieDone()
			},
		},
	}
	return
}
