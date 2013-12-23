# OVH CLI

## How to use

#### 1 - Download executable

* Windows : ftp://ftp.toorop.fr/softs/ovh_cli/windows/ovh.exe
* MacOs : ftp://ftp.toorop.fr/softs/ovh_cli/macos/ovh
* Linux : ftp://ftp.toorop.fr/softs/ovh_cli/linux/ovh


#### 2 - Get a consumer key
In order to access to your account, the app need your authorization. You have to get "a consumer key" and put "Keep calm and carry on", the app di all the job for you, just run it and follow instructions :

On Linux and MacOs run app with :

	./ovh
	
On windows with :

	./ovh.exe
	

## Avalaible commands
We will consider Linux|MacOs version, just replace *ovh* by *ovh.exe* if you are using Windows.

All WORDS in uppercase are variables, words in lower cases are parts of the command to be executed.
  
### IP
#### List IP Block
	./ovh ip list
Will return all your IP
You can provide a third argument defining the type of IP returned. For exemple, if you only want IP attached to tour dedicated server, run the command :

	./ovh ip list dedicated
	
Available type are :

* cdn
* dedicated
* hosted_ssl
* loadBalancing
* mail
* pcc
* pci
* vpn
* vps
* xdsl

 
 
### FIREWALL
All commands concerning firewall start with :

	./ovh ip fw
	
#### List IPs of an IP block which are under firewall

	./ovh ip fw IPBLOCK list
	
Where :

* IPBLOCK : an ip block given by "ovh ip list"

Response : Return a list of IPV4, one per line. Or error.	

Example :
	
	./ovh ip fw 176.31.189.121/32 list
	176.31.189.121	
	
#### Add an IP on firewall

	./ovh ip fw IPBLOCK IPV4 add
	
Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	

Response : "IPV4 added to firewall" if the command succeed an error otherwise.
	
Example :

	./ovh ip fw 176.31.189.121/32 176.31.189.121 add
	176.31.189.121 added to firewall	

#### Remove an IP from firewall

	./ovh ip fw IPBLOCK IPV4 remove

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	

Response : "IPV4 removed from firewall" if the command succeed an error otherwise.	
	
Example :
	
	./ovh ip fw 176.31.189.121/32 176.31.189.121 remove
	176.31.189.121 removed from firewall
		
#### Get Properties of a firewalled IP
	
	./ovh ip fw IPBLOCK IPV4 prop

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK		

Response : Properties on success, one per line. Error otherwise.

Example 

	./ovh ip fw 176.31.189.121/32 176.31.189.121 prop
	ipOnFirewall: 176.31.189.121
	Enabled: false
	State: ok				

#### Enable firewall

	./ovh ip fw IPBLOCK IPV4 enable

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK		

Response : "ok" if the command succeed. An error if not.

Example :
	
	./ovh ip fw 176.31.189.121/32 176.31.189.121 enable
	ok


#### Disable firewall

	./ovh ip fw IPBLOCK IPV4 disable

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK		

Response : "ok" if the command succeed. An error if not.

Example :
	
	./ovh ip fw 176.31.189.121/32 176.31.189.121 disable
	ok
	
#### Add a firewall rule

	 ./ovh ip fw IPBLOCK IPV4 addRule 'RULE'
	 

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	
* RULE : a JSON encoded rule object. See below.

##### JSON encoded rule format definition

Properties of the rule object are (* = requiered) :

* action* : Action of this rule. "permit" ot "deny".
* protocol* : "icmp" or "ipv4" or "tcp" or "udp"
* sequence* : sequence number of the rule (rule are excecuted from sequence=0 to sequence=n)
* source : Source ip for the rule. Any if not set.
* sourcePort : Source port range for the rule. Only with TCP/UDP protocol. It's an object with two properties :
	* from : first port
	* to : last port
* destinationPort : Destination port range for the rule. Only with TCP/UDP protocol. It's an object with two properties :
	* from : first port
	* to : last port
* tcpOption : It's a object with muliple boolean properties, if a propertie is set to true, the flag will be enabled :
	* ack
	* established
	* fin
	* psh
	* rst
	* syn
	* urg
* udpOption	: It's a object with muliple boolean properties, if a propertie is set to true, the flag will be enabled :
	* fragments 
	 

Examples :	 
	 
	  ./ovh ip fw 176.31.189.121/32 176.31.189.121 addRule '{"action":"deny","protocol":"udp","sequence":"0"}'
	  
Will add a rule	which deny all incoming udp traffic 

	  ./ovh ip fw 176.31.189.121/32 176.31.189.121 addRule '{"action": "permit", "destinationPort": {"from": 22, "to": 22},"protocol":"tcp","sequence": "1","source": "46.105.152.56/32"}'

Will add a rule which allow connection from IP 46.105.152.56 to port 22 (SSH)

	./ovh ip fw 176.31.189.121/32 176.31.189.121 addRule '{"action": "deny","destinationPort": {"to": 22,"from": 22},"protocol": "tcp","sequence": "2"}'
	
Will add a rule wich deny any connection to port 22 (SSH).

Rules are tested from sequence 0 to sequence n. When a rule matches it is applied and no other rules are tested. That mean with those examples that only IP 46.105.152.56 will be able to connect thru SSH to IP 176.31.189.121.
	  	
	
#### Remove a firewall rule

	 ./ovh ip fw IPBLOCK IPV4 remRule SEQUENCE

With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	
* SEQUENCE : Seqeunce number of the rule

Response : "Rule SEQUENE removed" on success or error.	 
Example :

	./ovh ip fw 176.31.189.121/32 176.31.189.121 remRule 1
	Rule 1 removed	

	
#### Get info about a firewall rule
	
	./ovh ip fw IPBLOCK IPV4 getRule SEQUENCE
	
With :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	
* SEQUENCE : Seqeunce number of the rule

Response : Formatted rule (see example) or error.

Example :

	./ovh ip fw 176.31.189.121/32 176.31.189.121 getRule 1
	Protocol: tcp
	Source: 8.8.8.8/32
	DestinationPort: eq 25
	Sequence: 1
	Options: urg psh ack syn fin rst
	Destination: 176.31.189.121/32
	Rule: permit tcp 8.8.8.8/32 range 10 20 176.31.189.121/32 eq 25 urg psh ack syn fin rst
	SourcePort: range 10 20
	State: ok
	CreationDate: 2013-12-20T17:45:07+01:00
	Action: permit
	
		


	


