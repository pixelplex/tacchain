package liquidstake_test

import (
	"testing"
	
	"cosmossdk.io/errors"
)

func TestErrorCodeConflict(t *testing.T) {
	// This test intentionally tries to register the same error code twice 
	// but with the same module name
	
	// Setup panic recovery
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected panic when registering duplicate error code, but no panic occurred")
		}
		t.Logf("Got expected panic: %v", r)
	}()
	
	// Register error code 2 with "invalid proof" message (similar to IBC)
	_ = errors.RegisterWithGRPCCode("commitment", 2, 2, "invalid proof")
	
	// Try to register the same error code again with the same module name
	// This should cause a panic
	_ = errors.RegisterWithGRPCCode("commitment", 2, 2, "another message")
}
