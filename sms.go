package main

import (
	"github.com/codegangsta/cli"
	"github.com/toorop/govh"
	"github.com/toorop/govh/sms"
)

// getFwCmds return commands for firewall subsection
func getSmsCmds(client *govh.OVHClient) (smsCmds []cli.Command) {
	sr, err := sms.New(client)
	if err != nil {
		return
	}
	smsCmds = []cli.Command{
		{
			Name:        "services",
			Usage:       "Return a list of sms services ",
			Description: "ovh sms services" + NLTAB + "Example: ovh sms services",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				services, err := sr.GetServices()
				dieOnError(err)
				println(formatOutput(services, c.Bool("json")))
				dieOk()
			},
		},
		{
			Name:        "send",
			Usage:       "Send an new SMS",
			Description: `ovh sms send SERVICE --from SENDER --to RECEIVER1 [RECEIVER2 RECEIVER3...RECEIVERX] --message MESSAGE + NLTAB + "Example: ovh sms send sms-st2-1 --from +33979XXXX --to +336222XXXX +336221XXXX --message "Test from ovh-cli"`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "from", Value: "", Usage: "The sender phone number in international format (+33XXXXXX for France for ex). Required."},
				cli.StringSliceFlag{Name: "to", Value: &cli.StringSlice{}, Usage: "Receiver phone number. If you have multiple receivers add on --receiver flag by reciever. Requiered."},
				cli.StringFlag{Name: "message", Value: "", Usage: "The message you want to send. Required"},
				cli.BoolFlag{Name: "json", Usage: "output as JSON"},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)

				// sender
				sender := c.String("from")
				if sender == "" {
					dieBadArgs("Required flag --from is missing")
				}
				// recievers
				receivers := c.StringSlice("to")
				if receivers == nil {
					dieBadArgs("Required flag --to is missing")
				}

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
				dieOnError(err)
				println(formatOutput(resp, c.Bool("json")))
				dieOk()
			},
		},
	}
	return
}
