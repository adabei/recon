# Recon
Recon is a REPL-like RCON client extracted from the goldenbot project.

## Installation
Installation is simple if you have Go installed. Just run
`go get github.com/adabei/recon`
and it will create an executable in $GOPATH/bin.


## Usage
On Windows you would start recon like this:

`./recon.exe 127.0.0.1:28960 -p q3`

After it has started up (and not bugged you with errors) you can send RCON commands to the server by typing them 
and confirming with enter. _Note: you don't need to write "rcon" or your password_

*If* there is something to return, Recon will print it.

### Example

`>say "this sure looks easy"`

## License
Recon is released under the MIT license. See LICENSE for details.
