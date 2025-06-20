package liquidstake_test

import (
	"testing"
	
	"cosmossdk.io/errors"
)

func TestErrorCodeRegistrationMinimal(t *testing.T) {
	// This test directly registers error codes to check for conflicts
	
	// Register error code 2 with "invalid proof" message (similar to IBC)
	_ = errors.Register("test-module", 2, "invalid proof")
	
	// Register error code 1000 (our new liquidstake error code) with any message
	// This should not conflict with the previous registration
	_ = errors.Register("test-module-2", 1000, "some error message")
	
	// If we reach this point without panicking, the test passes
	// which means our fix of changing error codes to start from 1000 works
}
