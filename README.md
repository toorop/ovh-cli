# OVH API CLI

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
	ovh ip list
Will return all your IP
You can provide a third argument defining the type of IP returned. For exemple, if you only want IP attached to tour dedicated server, run the command :

	ovh ip list dedicated
	
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

	ovh ip fw
	
#### List IPs of an IP block which are under firewall

	ovh ip fw IPBLOCK list
	
Where :

* IPBLOCK : an ip block given by "ovh ip list"

Response : Return a list of IPV4, one per line. Or error.	

Example :
	
	ovh ip fw 176.31.189.121/32 list
	176.31.189.121	
	
#### Add an IP on firewall

	ovh ip fw IPBLOCK IPV4 add
	
Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	

Response : "IPV4 added to firewall" if the command succeed an error otherwise.
	
Example :

	ovh ip fw 176.31.189.121/32 176.31.189.121 add
	176.31.189.121 added to firewall	

#### Remove an IP from firewall

	ovh ip fw IPBLOCK IPV4 remove

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK	

Response : "IPV4 removed from firewall" if the command succeed an error otherwise.	
	
Example :
	
	ovh ip fw 176.31.189.121/32 176.31.189.121 remove
	176.31.189.121 removed from firewall
		
#### Get Properties of a firewalled IP
	
	ip fw IPBLOCK IPV4 prop

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

	ip fw IPBLOCK IPV4 enable

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK		

Response : "ok" if the command succeed. An error if not.

Example :
	
	./ovh ip fw 176.31.189.121/32 176.31.189.121 enable
	ok


#### Disable firewall

	ip fw IPBLOCK IPV4 disable

Where :

* IPBLOCK : an ip block given by "ovh ip list"
* IPV4 : an IP v4 from IPBLOCK		

Response : "ok" if the command succeed. An error if not.

Example :
	
	./ovh ip fw 176.31.189.121/32 176.31.189.121 disable
	ok
	
#### Get a firewall rule
	
	ip fw IPBLOCK IPV4 getRule SEQUENCE
	
Where :

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
	
		


	


