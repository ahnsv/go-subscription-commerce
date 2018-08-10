PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE):/home/azer/code/golang # You can remove or change the path after last colon.
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will store the server process id when it's running on development mode
PID=/tmp/.$(PROJECTNAME)-api-server.pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent
