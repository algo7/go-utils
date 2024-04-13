// Package network provides network-related utility functions.
package network

import (
	"fmt"
	"net"
)

// IsPortInUse checks if a given TCP port is already or lower than 1024 (privilege ports) in use on the local machine.
// It attempts to listen on the specified port and returns true if an error occurs
// (indicating the port is likely in use). If the port is available, it returns false.
// The port to check is passed as an integer argument.
//
// Example usage:
//
//	port := 8500
//	if IsPortInUse(port) {
//	    fmt.Printf("Port %d is not available or lower than 1024.\n", port)
//	} else {
//	    fmt.Printf("Port %d is available.\n", port)
//	}
func IsPortInUse(port int) bool {

	// Check if the port is a privilege port (lower than 1024) or invalid port (higher than 65535).
	if port < 1024 || port > 65535 {
		return true
	}

	// Attempt to listen on the specified port.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// If there is an error, it likely means the port is in use.
		return true
	}

	// Close the listener and return false (the port is not in use).
	defer listener.Close()
	return false
}
