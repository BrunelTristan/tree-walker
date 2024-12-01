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
