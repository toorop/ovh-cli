package main

import (
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	"github.com/codegangsta/cli"
	"strconv"
	"strings"
)

// getFwCmds return commands for firewall subsection
func getFwCmds(client *govh.OvhClient) (fwCmds []cli.Command) {
	ipr, err := ip.New(client)
	if err != nil {
		return
	}

	fwCmds = []cli.Command{
		{
			Name:        "list",
			Usage:       "List IPs, of a given block, that are under firewall.",
			Description: "ovh fw list IPBLOCK" + NLTAB + "Example: ovh fw list 91.121.228.135/32 ",

			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					dieBadArgs()
				}
				ips, err := ipr.FwListIpOfBlock(ip.IpBlock{c.Args().First(), ""})
				handleErrFromOvh(err)
				for _, ip := range ips {
					fmt.Println(ip)
				}
				dieOk()
			},
		},
		{
			Name:        "add",
			Usage:       "Add an IP of IPBLOCK on firewall.",
			Description: "ovh fw add IPBLOCK IP" + NLTAB + "Example: ovh fw add 92.222.14.249/32 92.222.14.249",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				err := ipr.FwAddIp(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1))
				handleErrFromOvh(err)
				dieDone()
			},
		},
		{
			Name:        "remove",
			Usage:       "Remove an IP of IPBLOCK from firewall.",
			Description: "ovh fw remove IPBLOCK IP" + NLTAB + "Example: ovh fw remove 92.222.14.249/32 92.222.14.249",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				err := ipr.FwRemoveIp(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1))
				handleErrFromOvh(err)
				dieDone()
			},
		},
		{
			Name:        "getProperties",
			Usage:       "Get properties of an IP on the firewall.",
			Description: "ovh fw getProperties IPBLOCK IP " + NLTAB + "Example: ovh fw getProperties 92.222.14.249/32 92.222.14.249",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				p, err := ipr.FwGetIpProperties(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1))
				handleErrFromOvh(err)
				dieOk(fmt.Sprintf("Ip: %s%sEnabled: %t%sState: %s", p.Ip, NL, p.Enabled, NL, p.State))
			},
		},
		{
			Name:        "update",
			Usage:       "Update an IP on the firewall.",
			Description: "ovh fw update IPBLOCK IP [--flag...]" + NLTAB + "Example: ovh fw update 92.222.14.249/32 92.222.14.249 --enable true",
			Flags: []cli.Flag{
				cli.StringFlag{"enabled", "", "Set enabled state of the IP (true|false).", ""},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				fEnabled := c.Bool("enabled")
				err := ipr.FwUpdateIp(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1), fEnabled)
				handleErrFromOvh(err)
				dieDone()
			},
		},
		{
			Name:        "listRules",
			Usage:       "Return a list ao rule sequences.",
			Description: "ovh fw listRule IPBLOCK IP [--state]" + NLTAB + "Example: ovh fw listRule 92.222.14.249/32 92.222.14.249 --state ok",
			Flags: []cli.Flag{
				cli.StringFlag{"state", "", "Filter on state (creationPending|ok|removalPending).", ""},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				var sequences []int
				if c.IsSet("state") {
					fState := c.String("state")
					sequences, err = ipr.FwListRules(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1), fState)
				} else {
					sequences, err = ipr.FwListRules(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1))
				}
				handleErrFromOvh(err)
				for _, seq := range sequences {
					fmt.Println(seq)
				}
				dieOk()
			},
		},
		{
			Name:        "addRule",
			Usage:       "Add a new rule on an IP.",
			Description: "ovh fw addRule [--flag...] IPBLOCK IP  " + NLTAB + "Example: ovh fw addRule --action deny --protocol tcp --toPort 22 --sequence 0 92.222.14.249/32 92.222.14.249",
			Flags: []cli.Flag{
				cli.StringFlag{"action", "", "Action on this rule (deny|permit). Required.", ""},
				cli.StringFlag{"sequence", "", "Sequence number of your rule. Required.", ""},
				cli.StringFlag{"protocol", "", "Network protocol (ah|esp|gre|icmp|ipv4|tcp|udp). Requiered.", ""},
				cli.StringFlag{"fromPort", "", "Source port for your rule. Only with TCP/UDP protocol", ""},
				cli.StringFlag{"fromIp", "", "Source ip for your rule. Any if not set.", ""},
				cli.StringFlag{"toPort", "", "Destination port for your rule. Only with TCP/UDP protocol.", ""},
				cli.StringFlag{"tcpFragments", "", "Can only be used with TCP protocol (true|false)", ""},
				cli.StringFlag{"tcpOption", "", "Can only be used with TCP protocol (established|syn)", ""},
			},
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				rule := ip.FwRule2Add{}

				// action
				if !c.IsSet("action") {
					dieBadArgs()
				}
				action := strings.ToLower(c.String("action"))
				if !inSliceStr(action, []string{"deny", "permit"}) {
					dieBadArgs()
				}
				rule.Action = action

				// sequence
				if !c.IsSet("sequence") {
					dieBadArgs()
				}
				sequence := c.Int("sequence")
				rule.Sequence = sequence

				// protocol
				if !c.IsSet("protocol") {
					dieBadArgs()
				}
				protocol := strings.ToLower(c.String("protocol"))
				if !inSliceStr(protocol, []string{"ah", "esp", "gre", "icmp", "ipv4", "tcp", "udp"}) {
					dieBadArgs()
				}
				rule.Protocol = protocol

				// fromPort
				if c.IsSet("fromPort") {
					rule.FromPort = c.Int("fromPort")
				}

				// fromIP
				if c.IsSet("fromIp") {
					rule.FromIp = c.String("fromIp")
				}

				// toPort
				if c.IsSet("toPort") {
					rule.ToPort = c.Int("toPort")
				}

				// fwTcpOption
				fwTcpOption := ip.FwTcpOption{}
				flagFwTcpOption := false

				// tcpOptionFragment
				if c.IsSet("tcpFragments") {
					fwTcpOption.Fragments = c.Bool("tcpFragments")
					flagFwTcpOption = true
				}

				// tcpOption
				if c.IsSet("tcpOption") {
					tcpOption := c.String("tcpOption")
					if !inSliceStr(tcpOption, []string{"established", "syn"}) {
						dieBadArgs()
					}
					fwTcpOption.Option = tcpOption
					flagFwTcpOption = true

				}
				if flagFwTcpOption {
					rule.TcpOption = &fwTcpOption
				}
				handleErrFromOvh(ipr.FwAddRule(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1), rule))
				dieDone()
			},
		},
		{
			Name:        "removeRule",
			Usage:       "Remove a firwall rule.",
			Description: "ovh fw removeRule IPBLOCK IP SEQUENCE" + NLTAB + "Example: ovh fw removeRule 92.222.14.249/32 92.222.14.249 0",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 3)
				sequence, err := strconv.ParseInt(c.Args().Get(2), 10, 16)
				if err != nil {
					dieError(err)
				}
				handleErrFromOvh(ipr.FwRemoveRule(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1), int(sequence)))
				dieDone()
			},
		},
		{
			Name:        "getRuleProperties",
			Usage:       "Get properties of firewall rule.",
			Description: "ovh fw getRuleProperties IPBLOCK IP SEQUENCE " + NLTAB + "Example: ovh fw getRuleProperties 92.222.14.249/32 92.222.14.249 0",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 3)
				sequence, err := strconv.ParseInt(c.Args().Get(2), 10, 16)
				if err != nil {
					dieError(err)
				}
				p, err := ipr.FwGetRuleProperties(ip.IpBlock{c.Args().First(), ""}, c.Args().Get(1), int(sequence))
				handleErrFromOvh(err)
				dieOk(fmt.Sprintf("Sequence: %d%sCreated: %s%sProtocol: %s%sFromIp: %s%sFromPort: %s%sToIP: %s%sToPort: %s%sAction: %s%sRule:%s%sState: %s%sTcpOption: %s%sFragments: %t", p.Sequence, NL, p.CreationDate, NL, p.Protocol, NL, p.FromIp, NL, p.FromPort, NL, p.ToIp, NL, p.ToPort, NL, p.Action, NL, p.Rule, NL, p.State, NL, p.TcpOption, NL, p.Fragments))
			},
		},
	}
	return
}
