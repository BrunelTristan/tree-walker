package composition

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tree-walker/model/configuration"
)

func TestBuild(t *testing.T) {
	conf := &configuration.LaunchingConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()
}

func TestComposeWalker(t *testing.T) {
	conf := &configuration.LaunchingConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()

	walker := root.ComposeWalker()

	assert.NotNil(t, walker, "Should return built walker")
}

func TestComposeTree(t *testing.T) {
	conf := &configuration.LaunchingConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()

	christmasTree := root.ComposeTree()

	assert.NotNil(t, christmasTree, "Should return built tree")
}
