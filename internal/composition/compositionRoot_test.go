package composition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	root := NewCompositionRoot()

	root.Build()
}

func TestComposeWalker(t *testing.T) {
	root := NewCompositionRoot()

	root.Build()

	walker := root.ComposeWalker()

	assert.NotNil(t, walker, "Should return built walker")
}

func TestComposeTree(t *testing.T) {
	root := NewCompositionRoot()

	root.Build()

	christmasTree := root.ComposeTree()

	assert.NotNil(t, christmasTree, "Should return built tree")
}
