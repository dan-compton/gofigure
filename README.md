# gofigure
 
* Provides versioned, arbitrary configuration data (as limited by struct.proto supported types) 
* Provides GRPC endpoints for creating, updating, and checking configuration data
* Allows services to dynamic reconfigure themselves (like: let me update this log level without redeploying this service)
* Uses BoltDB as store


# Client usage

```
# Installation
go install github.com/dan-compton/gofigure/cmd/gfclient

# Create new service configuration
gfclient new -s someservice -v someconfigversion -c config_file.yml -a somegofigureserver:port

# Update remote service configuration
gfclient update -s someservice -v someconfigversion -c config_file.yml -a somegofigureserver:port

# Get remote service configuration
gfclient get -s someservice -v someconfigversion -o config_file.yml -a somegofigureserver:port

```

# TODO

* implementation of interface defined in gofigure package grpc
* remove request status from grpc
* store configuration state locally
