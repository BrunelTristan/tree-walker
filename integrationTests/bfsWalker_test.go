package integrationTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tree-walker/internal/composition"
	"tree-walker/model/tree"
)

func TestWalkOnWerySimpleTree(t *testing.T) {
	root := composition.NewCompositionRoot()

	root.Build()

	walker := root.ComposeWalker()

	rootNode := tree.Node{}
	firstNode := tree.Node{}
	secondNode := tree.Node{}
	lastNode := tree.Node{}

	unexploredTree := tree.Tree{
		Nodes: []tree.Node{
			rootNode,
			firstNode,
			secondNode,
			lastNode},
		Links: []tree.Link{
			tree.Link{Nodes: [2]*tree.Node{&rootNode, &firstNode}},
			tree.Link{Nodes: [2]*tree.Node{&rootNode, &secondNode}},
			tree.Link{Nodes: [2]*tree.Node{&secondNode, &lastNode}},
		}}

	path := walker.Walk(&unexploredTree, &rootNode, &lastNode)

	assert.NotEmpty(t, path, "Should found simple path")
	assert.Len(t, path.Nodes, 3, "Should have path with only three node")
	if 3 == len(path.Nodes) {
		assert.Equal(t, &rootNode, path.Nodes[0], "Should have path with root node")
		assert.Equal(t, &firstNode, path.Nodes[1], "Should have path with then first node")
		assert.Equal(t, &lastNode, path.Nodes[2], "Should have path ended by last node")
	}
}
