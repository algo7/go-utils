package network

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsPortInUse(t *testing.T) {

	assert := assert.New(t)

	// Test case 1: Port is a privilege port
	port := 80
	expected := true
	result := IsPortInUse(port)
	assert.Equal(expected, result, "Expected IsPortInUsed(%d) to be %v, but got %v", port, expected, result)

	// Test case 2: Port is not in use
	port = 8087
	expected = false
	assert.Equal(expected, IsPortInUse(port), "Expected IsPortInUsed(%d) to be %v, but got %v", port, expected, IsPortInUse(port))

	// Test case 3: Port is already in use
	// Make a channel to signal when the server is ready
	serverReady := make(chan struct{})

	// Make a channel to signal when the test is done
	testDone := make(chan struct{})

	// Start a dummy server to occupy the port
	port = 8888
	go func() {

		listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
		defer listener.Close()

		// Signal that the server is ready
		close(serverReady)

		<-testDone       // Wait for the test to be done
		listener.Close() // Then close the listener
	}()

	// Wait for the signal or timeout
	select {
	case <-serverReady:
		// Server is ready, proceed with the test
	case <-time.After(5 * time.Second):
		assert.Fail("Timeout waiting for server to start on port %d", port)
	}

	expected = true
	result = IsPortInUse(port)
	assert.Equal(expected, result, "Expected IsPortInUsed(%d) to be %v, but got %v", port, expected, result)

	close(testDone)
}
