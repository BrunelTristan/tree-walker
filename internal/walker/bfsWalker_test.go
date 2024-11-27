package walker

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tree-walker/model/tree"
)

func TestWalkOnEmptyTree(t *testing.T) {
	walker := BfsWalker{}

	unexploredTree := tree.Tree{}
	rootNode := tree.Node{}
	leafNode := tree.Node{}

	path := walker.Walk(unexploredTree, rootNode, leafNode)

	assert.Empty(t, path, "Should not found any path")
}
