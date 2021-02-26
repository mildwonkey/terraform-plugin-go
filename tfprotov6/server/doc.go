// Package tf6server implementations a server implementation for running
// tfprotov6.ProviderServers as gRPC servers.
//
// Providers will likely be calling tf6server.Serve from their main function to
// start the server so Terraform can connect to it.
package tf6server
