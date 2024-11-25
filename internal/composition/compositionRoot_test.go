package composition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	root := CompositionRoot{}

	root.Build()
}

func TestComposeWalker(t *testing.T) {
	root := CompositionRoot{}

	walker := root.ComposeWalker()

	assert.NotEmpty(t, walker, "Should return built walker")
}
