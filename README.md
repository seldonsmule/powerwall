# powerwall
This class is an example of how to use the APIs that are available from your Tesla powerwall/gateway system.  Done in Go

None of this would be possible without all the documentation provided here:
### https://github.com/vloschiavo/powerwall2

# Requirements
## Tesla Powerwall and gateway
## Know the IP address of the gateway (intenal webserver)
## Created a local login with password on the local gateway's web interface
## Cert file
The file "powerwall.cer" is the certificate for my powerwall.  It assumes you have a DNS entry of "powerwall" pointing to the IP address in your house.

## get_structs
# cmds
Example tools in the cmds folder.  

## get_structs
Gets the bulk of the useful json responses and creates structs
## guesstimate
Demonstrates how to use the generated structs to create a command.  This gets the basics status of a solar/powerwall system
