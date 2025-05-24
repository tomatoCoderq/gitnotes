package tools

import (
	"testing"
)

func TestResolveGitRef(t *testing.T) {
	ref, err := ResolveGitRef("9934f41b1bb25d75952b9b845b73a9fb1deb1497")
	if err != nil {
		t.Errorf("Error %v", err)
	}

	if ref != "9934f4" {
		t.Errorf("Got %v Want %v", ref, "9934f")
	}
}
