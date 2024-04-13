// Package bootstrap provides utilities for bootstrapping an application.
package bootstrap

import (
	"os"
	"strconv"
)

// GetEnvWithDefaultInt retrieves an environment variable as an int or returns a default value if not set or on parse error.
// Example usage:
//
//	port := GetEnvWithDefaultInt("PORT", 8080)
//	fmt.Printf("Using port: %d\n", port)
func GetEnvWithDefaultInt(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		// log.Printf("%s is not set, using default value: %d", key, defaultValue)
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		// log.Printf("Failed to parse %s as int: %v, using default value: %d", key, err, defaultValue)
		return defaultValue
	}
	return value
}

// GetEnvWithDefaultStr retrieves a string environment variable or returns a default value if not set.
//
// Example usage:
//
//	dir := GetEnvWithDefaultStr("DIR", "/tmp")
//	fmt.Printf("Using directory: %s\n", dir)
func GetEnvWithDefaultStr(key string, defaultValue string) string {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		// log.Printf("%s is not set, using default value: %s", key, defaultValue)
		return defaultValue
	}
	return valueStr
}
