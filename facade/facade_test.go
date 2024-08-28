package facade

import "testing"

var except = "A module running\nB module running"

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != except {
		t.Fatalf("except %s, return %s", except, ret)
	}
}
