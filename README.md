# OVH CLI


### WARNING reboot in progress

ovh-cli brings OVH services to the command line.


## How to use

#### 1 - Download executable

* [Windows](https://github.com/toorop/ovh-cli/releases/download/2.0.1/ovh_windows.exe)
* [MacOS](https://github.com/toorop/ovh-cli/releases/download/2.0.1/ovh_darwin)
* [Linux](https://github.com/toorop/ovh-cli/releases/download/2.0.1/ovh_linux)


#### 2 - Get a consumer key
In order to access to your account, the app need your authorization. You have to get "a consumer key" and record it as an environement variable.

"Keep calm and carry on", the app will do the job for you, just run it and follow instructions :

On Linux and MacOs run app with :

	./ovh

On windows with :

	./ovh.exe


## Avalaible commands
We will consider Linux (or MacOS) version, just replace *ovh* by *ovh.exe* if you are using Windows.

* WORDS in uppercase are variables
* words in lower cases are parts of the command to be executed.
* []: optionnal
* word...: means one or multiple

You can at any moment get help by adding the --help flag.

Example :

	$ ./ovh fw --help
	NAME:
   		ovh fw - Firewall commands

	USAGE:
   		ovh fw [subsection] command [command options] [arguments...]

	COMMANDS|SUBSECTION:
   		list			List IPs, of a given block, that are under firewall.
   		add				Add an IP of IPBLOCK on firewall.
   		remove			Remove an IP of IPBLOCK from firewall.
   		getProperties	Get properties of an IP on the firewall.
   		update			Update an IP on the firewall.
   		addRule			Add a new rule on an IP.
   		removeRule		Remove a firwall rule.
   		help, h			Shows a list of commands or help for one command

	OPTIONS:
   		--help, -h	show help


### Warning doc is outdated, use CLI help command instead (note that you can clone this repo, update the doc and do a PR ;))


* [IP](#ip)
	* [List IP blocks](#list-ip-blocks)
	* [Get IP block properties](#get-block-properties)
	* [Firewall](#firewall)
		* [List IP under firewall](#list-ips-of-an-ip-block-which-are-under-firewall)
		* [Add an IP on the firewall](#add-an-ip-on-firewall)
		* [Remove an IP from firewall](#remove-an-ip-from-firewall)
		* [Get properties of an IP](#get-properties-of-a-firewalled-ip)
		* [Update properties of an IP](#update-ip-properties)
		* [List firewall rules](#list-firewall-rules)
		* [Add a rule on an IP](#add-a-firewall-rule)
		* [Remove a firewall rule](#remove-a-firewall-rule)
		* [Get info about a firewall rule](#get-info-about-a-firewall-rule)
	* [Spam](#spam)
		* [List spamming IP of an IP block](#list-spamming-ip-of-an-ip-block)
  		* [Get info about a spamming IP](#get-more-info-about-a-spamming-ip)
  		* [Unblock a blocked IP](#unblock-a-blocked-ip)
  		* [Get statistics about a spamming IP for a time period](#get-statistics-about-a-spamming-ip-for-a-time-period)
  		* [Get IP which are currently blocked for spam](#get-ip-which-are-currently-blocked-for-spam-helper)
* [Server](#server)
	* [Get list of your dedicated servers](2#get-list-of-your-dedicated-servers)
	* [Get server properties](#get-server-properties)
	* [Reboot server](#reboot-server)
	* [Get server tasks](#get-server-tasks)
	* [Get server task properties](#get-server-task-properties)
	* [Cancel server task](#cancel-server-task)
* [SMS](#sms)
	* [List SMS services ](#list-sms-services)
	* [Send a SMS](#send-a-sms)
* [Cloud](#cloud)
	* [List cloud passports](#list-cloud-passports)
	* [List projects ID](#List-projects-id)
  	* [Get project info](#get-project-info)

## IP Block
#### List IP blocks
	./ovh ip block list [--flag...]
Will return your IP blocks

You can use filter flags :

* --desc: by descrition
* --ip: by IP
* --routedTo: by routing
* --type: by type (all|cdn|dedicated|failover|hosted_ssl|housing|loadBalancing|mail|pcc|pci|private|vps|vpn|vrack|xdsl)
* --json display output as JSON

Example: if you want IP blocks attached to dedicated server, run the command :

	./ovh ip block list --type dedicated

#### Get block properties
Return properties of an IP block

	./ovh ip block properties IPBLOCK

Options:

* --json display output as JSON

## FIREWALL
All commands concerning firewall start with :

	./ovh fw

#### List IPs of an IP block which are under firewall

	./ovh fw list IPBLOCK

Where :

* IPBLOCK : an ip block given by "ovh ip list"

Response : Return a list of IPV4, one per line. Or error.

Example :

	./ovh fw list 176.31.189.121/32
	176.31.189.121

#### Add an IP on firewall

	./ovh fw add IPBLOCK IPV4

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Response : "Done!" if the command succeed an error otherwise.

Example :

	./ovh fw 176.31.189.121/32 176.31.189.121 add
	Done!

#### Remove an IP from firewall

	./ovh fw remove IPBLOCK IPV4

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Response : "IPV4 removed from firewall" if the command succeed an error otherwise.

Example :

	./ovh ip fw 176.31.189.121/32 176.31.189.121 remove
	176.31.189.121 removed from firewall

#### Get Properties of a firewalled IP

	./ovh fw getProperties IPBLOCK IPV4

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Response : Properties on success, one per line. Error otherwise.

Example

	./ovh fw getProperties 176.31.189.121/32 176.31.189.121
	Ip: 176.31.189.121
	Enabled: false
	State: ok

#### Update IP properties

	./ovh fw update IPBLOCK IPV4 [--flag...]

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Available flags :

* --enabled bool: enabled|disabled IP on firewall


Response : "Done!" if the command succeed. An error if not.

Example :

	ovh fw update 92.222.14.249/32 92.222.14.249 --enabled true
	Done!

#### List firewall rules

	ovh fw listRules IPBLOCK IPV4 [--state creationPending|ok|removalPending]

With:

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Flags:

 * --state: State of the rule (creationPending|ok|removalPending)

Example:

	ovh fw listRules 92.222.14.249/32 92.222.14.249 --state ok
	1
	0

#### Add a firewall rule

	 ovh fw addRule IPBLOCK IPV4 [--flag...]

With:

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Flags:

 * --action: Action on this rule (deny|permit). Required.
 * --sequence: Sequence number of your rule. Required.
 * --protocol: Network protocol (ah|esp|gre|icmp|ipv4|tcp|udp). Requiered.
 * --fromPort: Source port for your rule. Only with TCP/UDP protocol.
 * --fromIp: Source ip for your rule. Any if not set.
 * -- toPort: Destination port for your rule. Only with TCP/UDP protocol.
 * --tcpFragment: Can only be used with TCP protocol (true|false).
 * --tcpOption: Can only be used with TCP protocol (established|syn)

Examples :
Add a rule	which deny all incoming udp traffic:

	  ovh fw addRule 92.222.14.249/32 92.222.14.249  --sequence 0 --action deny --protocole udp

Add a rule which allow connection from IP 46.105.152.56 to port 22 (SSH)

	ovh fw addRule 92.222.14.249/32 92.222.14.249 --sequence 0 --action permit --toPort 22 --fromIp 46.105.152.56/32 --protocol tcp  

Add a rule wich deny any connection to port 22 (SSH)

	ovh fw addRule 92.222.14.249/32 92.222.14.249 --sequence 1 --action deny --toPort 22 --protocol tcp

Will add a rule wich deny any connection to port 22 (SSH).

Rules are tested from sequence 0 to sequence n. When a rule matches it is applied and no other rules are tested. That mean with those examples that only IP 46.105.152.56 will be able to connect thru SSH to IP 92.222.14.249.


#### Remove a firewall rule

	 ovh fw removeRule IPBLOCK IPV4 SEQUENCE

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK
* SEQUENCE : Seqeunce number of the rule

Example :

	ovh fw removeRule 92.222.14.249/32 92.222.14.249 1
	Done!


#### Get info about a firewall rule

	ovh fw getRuleProperties IPBLOCK IPV4 SEQUENCE

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK
* SEQUENCE : Sequence number of the rule

Response : Formatted rule (see example) or error.

Example :

	ovh fw getRuleProperties 92.222.14.249/32 92.222.14.249 1
	Sequence: 1
	Created: 2014-06-02 07:49:42 +0000 UTC
	Protocol: tcp
	FromIp: 46.105.152.55/32
	FromPort:
	ToIP: 92.222.14.249/32
	ToPort: eq 23
	Action: permit
	Rule:permit tcp 46.105.152.55/32 92.222.14.249/32 eq 23 syn
	State: ok
	TcpOption: syn
	Fragments: false




### SPAM
#### List spamming IP of an IP block

	ovh spam listIp IPBLOCK [--state (blockedForSpam|unblocked|unblocking)]

With :

* IPBLOCK : an ip block given by "ovh ip list"

Flags:

* --state (optional) : current state of IP to list
	* blockeForSpam : Currently blocked for spam
	* unblocked : previously blocked
	* unblocking : an unblocking quiery have been made

Response : A list of IPv4 (one per line) or error.

Example :

	$ ovh spam listIp 178.33.223.32/28 --state unblocked
	178.33.223.41
	178.33.223.42
	178.33.223.43



#### Get more info about a spamming IP

	ovh spam getProperties IPBLOCK IPV4

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Response : Formated info (see example) or error

Example :

	ovh spam getProperties 91.121.228.135/32 91.121.228.135
	Blocked since (duration sec):0
	Last time: 2013-12-28T21:51:37+01:00
	IP: 91.121.228.135
	State: unblocked



#### Unblock a blocked IP

	ovh spam unblock IPBLOCK IPV4

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Response : "Done!" on success or error.

Exemple :

	ovh spam unblock 178.33.223.32/28 178.33.223.42
	ok


#### Get statistics about a spamming IP for a time period

	ovh spam getStats IPBLOCK IPV4 --from UNIX_TIMESTAMP_START --to UNIX_TIMESTAMP_STOP

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK

Flags

* --from: Unix timestamp representing the begining of the peiod (required).
* --to: Uni timestamp representing the end of the period (required).

Response : formated statistics (see example) or error

Example :

	ovh spam getStats 178.33.223.32/28 178.33.223.42 --from 1385251200 --to 1387882630
	Blocked for the last time: 12 Dec 13 22:52 +0100
	Number of emails sent: 141
	Number of spams sent: 124
	Average score: 521

	Detected Spams :


	Date: 12 Dec 13 20:49 +0100
	Message ID: 7847419050.QUJ8UNYQ543146@qcjeu.iiirdkehxxspb.va
	Destination IP: 91.121.228.128
	Score: 553

	Date: 12 Dec 13 20:49 +0100
	Message ID:
	Destination IP: 91.121.228.128
	Score: 303

	Date: 12 Dec 13 20:49 +0100
	Message ID:
	Destination IP: 91.121.228.128
	Score: 303

#### Get IP which are currently blocked for spam (helper)

	./ovh spam getBlocked

Response : List of blocked IP, one per line or error

Example :

	./ovh spam getBlocked
	178.33.223.41
	178.33.223.42
	178.33.223.43


## SERVER
### Get list of your dedicated servers

	./ovh server list

Response: all your server names. One per line.

Example :

	ovh server list
	ns309865.ovh.net
	ns205931.ovh.net
	ks323462.kimsufi.com
	ns212758.ovh.net
	ns26462.ovh.net

### Get server properties

	ovh server getProperties SERVER_NAME

With :

* SERVER_NAME : name of the server (from ./ovh server list)

Response : String with one property by line

Example :

	ovh server getProperties ks323462.kimsufi.com
	ID: 323462
	Name: ks323462.kimsufi.com
	Ip: 37.187.0.144
	Datacenter: gra1
	SupportLevel: pro
	ProfessionalUse: false
	CommercialRange: sk
	Os: ubuntu1304-server_64
	State: ok
	Reverse: ks323462.kimsufi.com.
	Monitored: true
	Rack: G102A03
	RootDevice:
	LinkSpeed: 100
	Bootid: 1


### Reboot server
	ovh server reboot SERVER_NAME

With :

* SERVER_NAME : name of the server (from ./ovh server list)

Response : A server task (see example)

Example :

	$ ovh server reboot ks323462.kimsufi.com
	Task ID: 2296596
	Function: hardReboot
	Status: init
	Comment: Reboot asked
	Last Upadte: 2014-04-01T08:43:23+02:00
	Start Date: 2014-04-01T08:43:23+02:00
	Done Date:


### Get server tasks

	ovh server getTasks SERVER_NAME --function  --status

With :

* SERVER_NAME : name of the server (from "ovh server list")

Flags :

* --function (optional) : filter by function. See [OVH doc](https://api.ovh.com/console/#/dedicated/server/%7BserviceName%7D/task#GET) for availables functions.)
* --status (optional) : filter by status. See [OVH doc](https://api.ovh.com/console/#/dedicated/server/%7BserviceName%7D/task#GET) for availables status.

Response : A list of task ID (1 per line)

Example :

	ovh server getTasks ns309865.ovh.net --function hardReboot --status done
	2296596
	2296597


### Get server task properties
	ovh server getTaskProperties SERVER_NAME TASK_ID

With :

* SERVER_NAME : name of the server (from ./ovh server list)
* TASK_ID : id of the task as returned by getTasks

Response : 	A server task as string (see example)

Example :

	ovh server getTaskProperties ks323462.kimsufi.com 4319579
	Task ID: 4319579
	Function: hardReboot
	Status: done
	Comment: Reboot asked
	Last Upadte: 2014-06-02 07:53:10 +0000 UTC
	Start Date: 2014-06-02 07:51:20 +0000 UTC
	Done Date: 2014-06-02 07:53:10 +0000 UTC


### Cancel server task

	ovh server cancelTask SERVER_NAME TASK_ID

With :

* SERVER_NAME : name of the server (from ./ovh server list)
* TASK_ID : id of the task as returned by getTasks

Response : "Done!" on success, error... on error

Examples :

	./ovh server cancelTask ks323462.kimsufi.com 2297463
	Done!

	./ovh server cancelTask ks323462.kimsufi.com 2297463
	Error!
	This task is already done


## SMS
#### List SMS services

	ovh sms listServices

Response : list of SMS services, one per line, or Error

Example :

	./ovh sms listServices
	sms-st2-1

#### Send a SMS

	ovh sms send SMS_SERVICE [--flags]

With :

* SMS_SERVICE : A sms service returned by "ovh sms listServices"

Flags:

* --message: message to send (string - required)
* --receiver : Receiver phone number. If you have multiple receivers add on --receiver flag by reciever (string - required)
* -- sender : sender (string - requiered)
* noStopClause : bool
* priority : hight|low|medium|veryLow (string)
* validityPeriod : The maximum time -in minute(s)- before the message is dropped. default 2880 (int)
* senderForResponse : (bool)
* charset : default UTF-8 (string)
* coding : defaut 7bit (string)
* differedPeriod : The time -in minute(s)- to wait before sending the message (int)
* receiversSlotId :The receivers document url link in csv format (string)
* class : flash|sim|phoneDisplay|toolkit (string)

Response :
* on success : string ""Done" + jobs details
* on failure : error


Example :

	ovh sms send sms-sd2545-1 --sender +339XXXXX --receiver +336XXXXX --message "Test from ovh-cli"

	Job ID: 23027686
	Invalid receivers:
	Valid receivers: +33622203197
	Credits removed: 1
	Done!


## Cloud
#### List cloud passports

	ovh cloud getPassports

Response : list of cloud passports, one per line, or Error

Example :

	ovh cloud getPassports
	publiccloud-passport-xxxxx

#### List projects ID

	ovh cloud getProjectsId

Response : list of cloud projects id, one per line, or Error

Example :

	ovh cloud getProjectsId
	a402414c9977XXXXXXXXXXXXXXXXXXXX
	a402414c997XXXXXXXXXXXXXXXXXXXXX

### Get project info

	ovh cloud getProject PROJECT_ID

Response: info about this project

Example:

	ovh cloud getProject a402414c997746e396456d0068a3XXX
	Project Id: a402414c997746e396456d0068a3XXXX
	Status: ok
	Creation date: 2015-05-06T20:20:26+02:00
	Description:





## Roadmap

* /cloud
* /cdn/dedicated todo
* /cdn/website todo
* /cdn/webstorage todo
* /cloud todo
* /dedicated/installationTemplate todo
* /dedicated/nas todo
* /dedicated/nasha todo
* /dedicated/server partially done
* /dedicatedCloud todo
* /domain todo
* /email/exchange todo
* /email/exchange todo
* /hosting/web todo
* /ip partially done
* /license/cpanel todo
* /license/directadmin todo
* /license/plesk todo
* /license/virtuozzo todo
* /license/windows todo
* /license/worklight todo
* /me todo
* /newAccount todo
* /order todo
* /price todo
* /sms partially done
* /telephony todo
* /vps todo
* /vrack todo
* /xdsl todo
