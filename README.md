# Zinn

### Developer Notes

This project is my playground. And has been inspired by the cf cli project. 
There are many patterns I have used from that project included the flag parsing
library from `jessevdk`

As a cli, this project has a few packages, that *should* have specific
responsibilities. 

The entry point to the program is in `cmd/zinn/main.go`. The main function is 
responsible for creating a new parser from the go-flags package, parsing the 
command line arguments, and then calling the correct command. 
Most of `main.go` is very similar to the main.go in the cf cli, most notably
the way I have extended the Command interface from `go-flags`. This makes it
easier to inject the UI object, and setup http clients for the commands. 

The command package is where the Execute functions for each supported command
live. The pattern is this, each command is represented by a command struct
which supports Setup and Execute functions. The Setup function sets the UI for
the command to use and the http client for the command to use. The Execute
function will typically call a function on the provided http client, and take
the result of that http call and process it. The processed response is then
output to the UI. 

The api package is where all the logic for interacting with the Guild Wars 2
api lives. It includes a client package which handles http operations. The 
client should know how to make requests to the gw2 api, and unmarshal the 
responses

The UI package, should be responsible for formatting and printing data to the
screen in an easy to read way. It does this by providing an Out writer and an
Err writer (see io package for more info). Initially the public functions are
"DisplayError" and "DisplayText" but it could one day contain things such as
"DisplayTable", or even a way to read from stdin.

To run all the tests except integration, use the `bin/test` script. The 
integration tests require an internet connection, as they run the cli in anger




