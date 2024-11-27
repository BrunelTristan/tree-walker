package walker

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tree-walker/model/tree"
)

func TestWalkOnEmptyTree(t *testing.T) {
	walker := BfsWalker{}

	unexploredTree := tree.Tree{}

	path := walker.Walk(&unexploredTree, nil, nil)

	assert.Empty(t, path, "Should not found any path")
}

func TestWalkOnMonoNodeTree(t *testing.T) {
	walker := BfsWalker{}

	rootNode := tree.Node{}

	unexploredTree := tree.Tree{Nodes: []tree.Node{rootNode}}

	path := walker.Walk(&unexploredTree, &rootNode, &rootNode)

	assert.NotEmpty(t, path, "Should found simple path")
	assert.Len(t, path.Nodes, 1, "Should have path with only one node")
	assert.Equal(t, &rootNode, path.Nodes[0], "Should have path with onlyroot node")
}
