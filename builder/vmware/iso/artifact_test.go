package iso

import (
	"github.com/aaronbareford/packer/packer"
	"testing"
)

func TestArtifact_Impl(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packer.Artifact); !ok {
		t.Fatal("Artifact must be a proper artifact")
	}
}
