package function

import "testing"

// TestHandler will fail
func TestHandler(t *testing.T) {
	t.Fatalf("It didn't work.")
}
