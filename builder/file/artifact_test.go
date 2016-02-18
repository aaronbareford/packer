package file

import (
	"testing"

	"github.com/aaronbareford/packer/packer"
)

func TestNullArtifact(t *testing.T) {
	var _ packer.Artifact = new(FileArtifact)
}
