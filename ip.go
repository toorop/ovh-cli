package main

import (
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	"strings"
)

func ipHandler(cmd *Cmd) (resp string, err error) {
	// New govh client
	client := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)
	// New ip ressource
	ipr, err := ip.New(client)

	//debug(cmd.Action)

	switch cmd.Action {
	// List
	case "list":
		ipType := "all"
		if len(cmd.Args) > 2 {
			ipType = cmd.Args[2]
		}
		ips, err := ipr.List(ipType)
		if err != nil {
			dieError(err)
		}
		for _, i := range ips {
			resp = fmt.Sprintf("%s%s\r\n", resp, i.IP)
		}
		if len(resp) > 2 {
			resp = resp[0 : len(resp)-2]
		}
		break
	case "lb":
		if len(cmd.Args) < 3 {
			dieError("\"ip lb\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
		}
		var t []byte
		t, err = ipr.LbList()
		resp = string(t)
		break

	case "fw":
		// ip fw ipBlock.IP list
		// ip fw x.x.x.x/y
		// Return IP V4 list of this block which is under firewall
		if len(cmd.Args) == 4 && cmd.Args[3] == "list" {
			block := ip.IpBlock{cmd.Args[2], ""}
			ips, err := ipr.FwListIpOfBlock(block)
			if err != nil {
				dieError(err)
			}
			for _, i := range ips {
				resp = fmt.Sprintf("%s%s\r\n", resp, i)
			}
			if len(resp) > 2 {
				resp = resp[0 : len(resp)-2]
			}
			break
		}

		// Add IP to firewall
		// cmd : ip fw ibBlock.IP ipV4 add
		if len(cmd.Args) == 5 && cmd.Args[4] == "add" {
			block := ip.IpBlock{cmd.Args[2], ""}
			if err = ipr.FwAddIp(block, cmd.Args[3]); err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("%s added to firewall", cmd.Args[3]))
		}

		// Get properties of a firewalled IP
		// ip fw ipBlock.IP ipV4 prop
		if len(cmd.Args) == 5 && cmd.Args[4] == "prop" {
			block := ip.IpBlock{cmd.Args[2], ""}
			i, err := ipr.FwGetIpProperties(block, cmd.Args[3])
			if err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("ipOnFirewall: %s%sEnabled: %t%sState: %s", i.IpOnFirewall, NL, i.Enabled, NL, i.State))
			break
		}

		// Enable firewalll for IP ipv4
		// ip fw ipVlock ipV4 enable
		if len(cmd.Args) == 5 && cmd.Args[4] == "enable" {
			block := ip.IpBlock{cmd.Args[2], ""}
			err := ipr.FwSetFirewallEnable(block, cmd.Args[3], true)
			if err != nil {
				dieError(err)
			}
			dieOk("ok")
			break
		}

		// Disable firewalll for IP ipv4
		// ip fw ipVlock ipV4 disable
		if len(cmd.Args) == 5 && cmd.Args[4] == "disable" {
			block := ip.IpBlock{cmd.Args[2], ""}
			err := ipr.FwSetFirewallEnable(block, cmd.Args[3], false)
			if err != nil {
				dieError(err)
			}
			dieOk("ok")
			break
		}

		// Remove IPv4 from firewall
		// cmd : ip fw ipBlock.IP ipV4 remove
		if len(cmd.Args) == 5 && cmd.Args[4] == "remove" {
			block := ip.IpBlock{cmd.Args[2], ""}
			if err = ipr.FwRemoveIp(block, cmd.Args[3]); err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("%s removed from firewall", cmd.Args[3]))
		}

		// ip fw ipBlock.IP ipV4 listRules all

		// ip fw ipBlock.IP ipV4 addRules rule (as Json)

		// Get rule
		// ip fw ipBlock.IP ipV4 getRule sequence
		if len(cmd.Args) == 6 && cmd.Args[4] == "getRule" {
			block := ip.IpBlock{cmd.Args[2], ""}
			rule, err := ipr.FwGetRule(block, cmd.Args[3], cmd.Args[5])
			if err != nil {
				dieError(err)
			}
			out := ""
			if len(rule.Protocol) > 0 {
				out = fmt.Sprintf("%sProtocol: %s%s", out, rule.Protocol, NL)
			}
			if len(rule.Source) > 0 {
				out = fmt.Sprintf("%sSource: %s%s", out, rule.Source, NL)
			}
			if len(rule.DestinationPort) > 0 {
				out = fmt.Sprintf("%sDestinationPort: %s%s", out, rule.DestinationPort, NL)
			}

			out = fmt.Sprintf("%sSequence: %d%s", out, rule.Sequence, NL)

			if len(rule.Options) > 0 {
				out = fmt.Sprintf("%sOptions: %s%s", out, strings.Join(rule.Options, " "), NL)
			}
			if len(rule.Destination) > 0 {
				out = fmt.Sprintf("%sDestination: %s%s", out, rule.Destination, NL)
			}
			if len(rule.Rule) > 0 {
				out = fmt.Sprintf("%sRule: %s%s", out, rule.Rule, NL)
			}
			if len(rule.SourcePort) > 0 {
				out = fmt.Sprintf("%sSourcePort: %s%s", out, rule.SourcePort, NL)
			}
			if len(rule.State) > 0 {
				out = fmt.Sprintf("%sState: %s%s", out, rule.State, NL)
			}
			if len(rule.CreationDate) > 0 {
				out = fmt.Sprintf("%sCreationDate: %s%s", out, rule.CreationDate, NL)
			}
			if len(rule.Action) > 0 {
				out = fmt.Sprintf("%sAction: %s%s", out, rule.Action, NL)
			}
			dieOk(out[0 : len(out)-2])

		}

		// ip fw ipBlock.IP ipV4 delRule sequence

		break
	default:
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
	}
	return

}
