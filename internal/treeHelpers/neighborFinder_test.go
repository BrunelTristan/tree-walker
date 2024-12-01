package treeHelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tree-walker/model/tree"
)

func TestGetNeighborsWithNilNode(t *testing.T) {
	baseTree := tree.Tree{
		Links: []tree.Link{},
	}

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(&baseTree, nil)

	assert.Empty(t, neighbors, "Should not find any neighbor")
}

func TestGetNeighborsWithNilTree(t *testing.T) {
	someNode := tree.Node{}

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(nil, &someNode)

	assert.Empty(t, neighbors, "Should not find any neighbor")
}

func TestGetNeighborsWithoutLink(t *testing.T) {
	baseTree := tree.Tree{
		Links: []tree.Link{},
	}
	someNode := tree.Node{}

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(&baseTree, &someNode)

	assert.Empty(t, neighbors, "Should not find any neighbor")
}

func TestGetNeighborsWithNoUsefulNode(t *testing.T) {
	baseTree := tree.Tree{
		Links: []tree.Link{
			tree.Link{Nodes: [2]*tree.Node{&tree.Node{ID: 4}, &tree.Node{ID: 8}}},
			tree.Link{Nodes: [2]*tree.Node{&tree.Node{ID: 3}, &tree.Node{ID: 6}}},
			tree.Link{Nodes: [2]*tree.Node{&tree.Node{ID: 3}, &tree.Node{ID: 4}}},
		},
	}
	someNode := tree.Node{ID: 0}

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(&baseTree, &someNode)

	assert.Empty(t, neighbors, "Should not find any neighbor")
}

func TestGetNeighborsWithJustOneLink(t *testing.T) {
	baseTree := tree.Tree{
		Nodes: []tree.Node{
			tree.Node{ID: 0},
			tree.Node{ID: 1},
			tree.Node{ID: 2},
			tree.Node{ID: 3},
			tree.Node{ID: 4},
			tree.Node{ID: 8},
		},
	}
	baseTree.Links = []tree.Link{
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[0], &baseTree.Nodes[5]}},
	}

	someNode := &baseTree.Nodes[0]
	otherNode := &baseTree.Nodes[5]

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(&baseTree, someNode)

	assert.NotEmpty(t, neighbors, "Should find a neighbor")
	assert.Len(t, neighbors, 1, "Should have one neighbor")
	if 1 == len(neighbors) {
		assert.Equal(t, otherNode, neighbors[0], "Should have the neighbor")
	}
}

func TestGetNeighborsWithJustOneLinkBackwards(t *testing.T) {
	baseTree := tree.Tree{
		Nodes: []tree.Node{
			tree.Node{ID: 0},
			tree.Node{ID: 1},
			tree.Node{ID: 2},
			tree.Node{ID: 3},
			tree.Node{ID: 4},
			tree.Node{ID: 8},
		},
	}
	baseTree.Links = []tree.Link{
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[5], &baseTree.Nodes[0]}},
	}

	someNode := &baseTree.Nodes[0]
	otherNode := &baseTree.Nodes[5]

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(&baseTree, someNode)

	assert.NotEmpty(t, neighbors, "Should find a neighbor")
	assert.Len(t, neighbors, 1, "Should have one neighbor")
	if 1 == len(neighbors) {
		assert.Equal(t, otherNode, neighbors[0], "Should have the neighbor")
	}
}

func TestGetNeighborsWithMultipleLinks(t *testing.T) {
	baseTree := tree.Tree{
		Nodes: []tree.Node{
			tree.Node{ID: 0},
			tree.Node{ID: 1},
			tree.Node{ID: 2},
			tree.Node{ID: 3},
			tree.Node{ID: 4},
			tree.Node{ID: 8},
		},
	}
	baseTree.Links = []tree.Link{
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[5], &baseTree.Nodes[0]}},
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[5], &baseTree.Nodes[4]}},
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[3], &baseTree.Nodes[2]}},
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[0], &baseTree.Nodes[1]}},
		tree.Link{Nodes: [2]*tree.Node{&baseTree.Nodes[0], &baseTree.Nodes[3]}},
	}

	someNode := &baseTree.Nodes[0]

	finder := NewNeighborFinder()

	neighbors := finder.GetNeighbors(&baseTree, someNode)

	assert.NotEmpty(t, neighbors, "Should find neighbors")
	assert.Len(t, neighbors, 3, "Should have three neighbors")
	if 3 == len(neighbors) {
		assert.Contains(t, neighbors, &baseTree.Nodes[1], "Should contains node 1")
		assert.Contains(t, neighbors, &baseTree.Nodes[3], "Should contains node 3")
		assert.Contains(t, neighbors, &baseTree.Nodes[5], "Should contains node 5")
	}
}
