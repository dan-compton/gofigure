# gofigure
 
* Provides versioned, arbitrary configuration data (as limited by struct.proto supported types) 
* Provides GRPC endpoints for creating, updating, and checking configuration data
* Allows services to dynamic reconfigure themselves (like: let me update this log level without redeploying this service)
* Uses BoltDB as store

# TODO

* implementation of interface defined in gofigure package grpc

