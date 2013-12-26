package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	//"os"
	"strconv"
	"strings"
	"time"
)

func ipHandler(cmd *Cmd) (resp string, err error) {
	// New govh client
	client := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)
	// New ip ressource
	ipr, err := ip.New(client)

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

		// Get rules sequences
		// ip fw ipBlock.IP ipV4 listRules all
		if len(cmd.Args) >= 5 && cmd.Args[4] == "listRules" {
			block := ip.IpBlock{cmd.Args[2], ""}
			state := ""
			if len(cmd.Args) == 6 {
				state = cmd.Args[6]
			}
			t, err := ipr.FwGetRulesSequences(block, cmd.Args[3], state)
			if err != nil {
				dieError(err)
			}
			var r string
			if len(t) > 0 {
				r = fmt.Sprintf("%d", t[0])
				for _, s := range t[1:] {
					r = fmt.Sprintf("%s%s%d", r, NL, s)
				}
			}
			dieOk(r)
		}

		// Add rule
		// ip fw ipBlock.IP ipV4 addRule rule (as Json)
		if len(cmd.Args) == 6 && cmd.Args[4] == "addRule" {
			block := ip.IpBlock{cmd.Args[2], ""}
			// Check json
			var rule ip.FirewallRule2Add
			err := json.Unmarshal([]byte(cmd.Args[5]), &rule)
			if err != nil {
				dieError("Rule error. See doc at : https://github.com/Toorop/ovh-cli", err)
			}
			err = ipr.FwAddRule(block, cmd.Args[3], rule)
			if err != nil {
				dieError(err)
			}
			dieOk("OK")
		}

		// Remove rule
		// ip fw ipBlock.IP ipV4 remRule ruleSequence
		if len(cmd.Args) == 6 && cmd.Args[4] == "remRule" {
			block := ip.IpBlock{cmd.Args[2], ""}
			sequence, err := strconv.Atoi(cmd.Args[5])
			if err != nil {
				dieError(err)
			}
			err = ipr.FwRemoveRule(block, cmd.Args[3], sequence)
			if err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("Rule %d removed", sequence))

		}

		// Get rule
		// ip fw ipBlock.IP ipV4 getRule sequence
		if len(cmd.Args) == 6 && cmd.Args[4] == "getRule" {
			block := ip.IpBlock{cmd.Args[2], ""}
			sequence, err := strconv.Atoi(cmd.Args[5])
			if err != nil {
				dieError(err)
			}
			rule, err := ipr.FwGetRule(block, cmd.Args[3], sequence)
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
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
		break

	case "spam":
		// List of spamming IP
		// ip spam ipBlock.IP listSpammingIp STATE
		if len(cmd.Args) >= 4 && cmd.Args[3] == "listSpammingIp" {
			block := ip.IpBlock{cmd.Args[2], ""}
			state := ""
			if len(cmd.Args) == 5 {
				state = cmd.Args[4]
			}
			ips, err := ipr.SpamGetSpammingIps(block, state)
			if err != nil {
				dieError(err)
			}
			for _, ip := range ips {
				fmt.Println(ip)
			}
			dieOk("")
		}

		// detailed info about a spamming IP
		// ip spam ipBlock.IP ipv4 details
		if len(cmd.Args) == 5 && cmd.Args[4] == "details" {
			block := ip.IpBlock{cmd.Args[2], ""}
			spamIp, err := ipr.SpamGetSpamIp(block, cmd.Args[3])
			if err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("Time: %d%sDate: %s%sIpSpamming: %s%sState: %s", spamIp.Time, NL, spamIp.Date, NL, spamIp.IpSpamming, NL, spamIp.State))
		}

		// Stats about a spamming IP
		// ip spam ipBlock.IP ipv4 stats FROM TO
		if len(cmd.Args) == 7 && cmd.Args[4] == "stats" {
			block := ip.IpBlock{cmd.Args[2], ""}

			from, err := strconv.ParseInt(cmd.Args[5], 10, 64)
			if err != nil {
				dieError(err)
			}

			to, err := strconv.ParseInt(cmd.Args[6], 10, 64)
			if err != nil {
				dieError(err)
			}
			spamStats, err := ipr.SpamGetIpStats(block, cmd.Args[3], time.Unix(from, 0), time.Unix(to, 0))
			if err != nil {
				dieError(err)
			}
			if spamStats == nil {
				dieOk("No spam stats for this period")
			}
			fmt.Printf("Blocked for the last time: %s%s", time.Unix(spamStats.Timestamp, 0).Format(time.RFC822Z), NL)
			fmt.Printf("Number of emails sent: %d%s", spamStats.Total, NL)
			fmt.Printf("Number of spams sent: %d%s", spamStats.NumberOfSpams, NL)
			fmt.Printf("Average score: %d%s%s", spamStats.AverageSpamScore, NL, NL)
			if len(spamStats.DetectedSpams) > 0 {
				fmt.Println("Detected Spams : ", NL)
			}
			for _, ds := range spamStats.DetectedSpams {
				fmt.Println("")
				fmt.Printf("\tDate: %s%s", time.Unix(ds.Date, 0).Format(time.RFC822Z), NL)
				fmt.Printf("\tMessage ID: %s%s", ds.MessageId, NL)
				fmt.Printf("\tDestination IP: %s%s", ds.DestinationIp, NL)
				fmt.Printf("\tScore: %d%s", ds.Spamscore, NL)
			}
			dieOk("")
		}

		// Unblock
		// ip spam ipBlock.IP ipv4 unblock
		if len(cmd.Args) == 5 && cmd.Args[4] == "unblock" {
			block := ip.IpBlock{cmd.Args[2], ""}
			err := ipr.SpamUnblockSpamIp(block, cmd.Args[3])
			if err != nil {
				dieError(err)
			}
			dieOk("ok")
		}

		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
		break
	case "getBlockedForSpam":
		// On va chercher les blocks
		ips, err := ipr.GetBlockedForSpam()
		if err != nil {
			dieError(err)
		}
		if len(ips) == 0 {
			dieOk("")
		}
		for _, i := range ips {
			fmt.Println(i)
		}
		dieOk("")

		// On les tests
		break

	default:
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
	}
	return

}
