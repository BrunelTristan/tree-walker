package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTreeWithoutNode(t *testing.T) {
	builder := NewBinaryTreeBuilder(0)

	tree := builder.BuildTree()

	assert.Nil(t, tree, "Should return no tree")
}

func TestBuildTreeVerySimple(t *testing.T) {
	nodeCount := 3
	builder := NewBinaryTreeBuilder(nodeCount)

	tree := builder.BuildTree()

	assert.NotNil(t, tree, "Should return built tree")
	assert.Len(t, tree.Nodes, nodeCount, "Should have node count")
	assert.Len(t, tree.Links, nodeCount-1, "Should connect every nodes")
	if len(tree.Links) == nodeCount-1 {
		assert.Equal(t, 0, tree.Links[0].Nodes[0].ID, "First link should start with root node")
		assert.Equal(t, 1, tree.Links[0].Nodes[1].ID, "First link should end with first node")
		assert.Equal(t, 0, tree.Links[1].Nodes[0].ID, "Second link should start with root node")
		assert.Equal(t, 2, tree.Links[1].Nodes[1].ID, "Second link should end with last node")
	}
}

func TestBuildTreeWithThreeLevel(t *testing.T) {
	nodeCount := 7
	builder := NewBinaryTreeBuilder(nodeCount)

	tree := builder.BuildTree()

	assert.NotNil(t, tree, "Should return built tree")
	assert.Len(t, tree.Nodes, nodeCount, "Should have node count")
	assert.Len(t, tree.Links, nodeCount-1, "Should connect every nodes")
	if len(tree.Links) == nodeCount-1 {
		assert.Equal(t, 0, tree.Links[0].Nodes[0].ID, "First link should start with root node")
		assert.Equal(t, 1, tree.Links[0].Nodes[1].ID, "First link should end with first node")
		assert.Equal(t, 0, tree.Links[1].Nodes[0].ID, "Second link should start with root node")
		assert.Equal(t, 2, tree.Links[1].Nodes[1].ID, "Second link should end with second node")
		assert.Equal(t, 1, tree.Links[2].Nodes[0].ID, "Third link should start with first node")
		assert.Equal(t, 3, tree.Links[2].Nodes[1].ID, "Third link should end with third node")
		assert.Equal(t, 1, tree.Links[3].Nodes[0].ID, "Fourth link should start with first node")
		assert.Equal(t, 4, tree.Links[3].Nodes[1].ID, "Fourth link should end with fourth node")
		assert.Equal(t, 2, tree.Links[4].Nodes[0].ID, "Fifth link should start with second node")
		assert.Equal(t, 5, tree.Links[4].Nodes[1].ID, "Fifth link should end with fifth node")
		assert.Equal(t, 2, tree.Links[5].Nodes[0].ID, "Sixth link should start with second node")
		assert.Equal(t, 6, tree.Links[5].Nodes[1].ID, "Sixth link should end with sixth node")
	}
}

func TestBuildTreeNotEquilibrium(t *testing.T) {
	nodeCount := 10
	builder := NewBinaryTreeBuilder(nodeCount)

	tree := builder.BuildTree()

	assert.NotNil(t, tree, "Should return built tree")
	assert.Len(t, tree.Nodes, nodeCount, "Should have node count")
	assert.Len(t, tree.Links, nodeCount-1, "Should connect every nodes")
}

func TestBuildTreeBigSequoia(t *testing.T) {
	nodeCount := 999
	builder := NewBinaryTreeBuilder(nodeCount)

	tree := builder.BuildTree()

	assert.NotNil(t, tree, "Should return built tree")
	assert.Len(t, tree.Nodes, nodeCount, "Should have node count")
	assert.Len(t, tree.Links, nodeCount-1, "Should connect every nodes")
}
