package bootstrap

import (
	"log"
	"os"
	"testing"
)

func TestGetEnvWithDefaultInt(t *testing.T) {
	// Prepare a logger to capture the output and prevent it from appearing during tests.
	log.SetOutput(os.NewFile(0, os.DevNull))

	// Test cases
	cases := []struct {
		name string
		key  string
		// Value to set for the environment variable. If empty, the variable will be unset.
		setValue     string
		defaultValue int
		expected     int
	}{
		{
			name:         "Variable set with valid int",
			key:          "TEST_VAR_VALID",
			setValue:     "10",
			defaultValue: 5,
			expected:     10,
		},
		{
			name:         "Variable set with invalid int",
			key:          "TEST_VAR_INVALID",
			setValue:     "invalid",
			defaultValue: 5,
			expected:     5,
		},
		{
			name:         "Variable not set",
			key:          "TEST_VAR_UNSET",
			setValue:     "",
			defaultValue: 5,
			expected:     5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Set or unset the environment variable as per the test case.
			if tc.setValue != "" {
				os.Setenv(tc.key, tc.setValue)
				// Unset the environment variable after the test.
				defer os.Unsetenv(tc.key)
			} else {
				os.Unsetenv(tc.key)
			}

			// Call the function under test.
			got := GetEnvWithDefaultInt(tc.key, tc.defaultValue)

			// Assert the result.
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}

func TestGetEnvWithDefaultStr(t *testing.T) {
	// Silence the logger during tests
	log.SetOutput(os.NewFile(0, os.DevNull))

	// Define test cases
	testCases := []struct {
		name         string
		key          string
		setValue     string
		defaultValue string
		expected     string
	}{
		{
			name:         "Environment variable set",
			key:          "TEST_ENV_VAR_SET",
			setValue:     "actual_value",
			defaultValue: "default_value",
			expected:     "actual_value",
		},
		{
			name:         "Environment variable not set",
			key:          "TEST_ENV_VAR_NOT_SET",
			setValue:     "",
			defaultValue: "default_value",
			expected:     "default_value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set or unset the environment variable based on the test case setup
			if tc.setValue != "" {
				os.Setenv(tc.key, tc.setValue)
				defer os.Unsetenv(tc.key) // Ensure clean-up
			} else {
				os.Unsetenv(tc.key)
			}

			// Call the function under test
			got := GetEnvWithDefaultStr(tc.key, tc.defaultValue)

			// Assert the expected outcome
			if got != tc.expected {
				t.Errorf("GetEnvWithDefaultStr(%q, %q) = %q, want %q", tc.key, tc.defaultValue, got, tc.expected)
			}
		})
	}
}
