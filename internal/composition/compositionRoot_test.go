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

	walker := root.composeWalker()

	assert.NotNil(t, walker, "Should return built walker")
}
