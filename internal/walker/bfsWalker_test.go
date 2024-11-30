package walker

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tree-walker/internal/generatedMocks"
	"tree-walker/model/tree"
)

func TestWalkOnEmptyTree(t *testing.T) {
	walker := NewBfsWalker(nil)

	unexploredTree := tree.Tree{}

	path := walker.Walk(&unexploredTree, nil, nil)

	assert.Empty(t, path, "Should not find any path")
}

func TestWalkOnMonoNodeTree(t *testing.T) {
	walker := NewBfsWalker(nil)

	unexploredTree := tree.Tree{Nodes: []tree.Node{tree.Node{}}}

	rootNode := &unexploredTree.Nodes[0]

	path := walker.Walk(&unexploredTree, rootNode, rootNode)

	assert.NotEmpty(t, path, "Should find simple path")
	assert.Len(t, path.Nodes, 1, "Should have path with only one node")
	assert.Equal(t, rootNode, path.Nodes[0], "Should have path with only root node")
}

func TestWalkWithoutExistingPath(t *testing.T) {
	mockController := gomock.NewController(t)

	neighborFinderMock := generatedMocks.NewMockINeighborFinder(mockController)

	walker := NewBfsWalker(neighborFinderMock)

	unexploredTree := tree.Tree{
		Nodes: []tree.Node{
			tree.Node{ID: 0},
			tree.Node{ID: 1},
			tree.Node{ID: 3},
			tree.Node{ID: 5},
		},
	}

	rootNode := &unexploredTree.Nodes[0]
	lastNode := &unexploredTree.Nodes[3]

	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, rootNode).
		Return([]*tree.Node{})

	path := walker.Walk(&unexploredTree, rootNode, lastNode)

	assert.Empty(t, path, "Should not find any path")
}

func TestWalkOnVerySimpleTree(t *testing.T) {
	mockController := gomock.NewController(t)

	neighborFinderMock := generatedMocks.NewMockINeighborFinder(mockController)

	walker := NewBfsWalker(neighborFinderMock)

	unexploredTree := tree.Tree{
		Nodes: []tree.Node{
			tree.Node{ID: 0},
			tree.Node{ID: 1},
			tree.Node{ID: 3},
			tree.Node{ID: 5},
		},
	}

	rootNode := &unexploredTree.Nodes[0]
	firstNode := &unexploredTree.Nodes[1]
	secondNode := &unexploredTree.Nodes[2]
	lastNode := &unexploredTree.Nodes[3]

	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, rootNode).
		Return([]*tree.Node{firstNode, secondNode})
	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, firstNode).
		Return([]*tree.Node{rootNode})
	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, secondNode).
		Return([]*tree.Node{rootNode, lastNode})

	path := walker.Walk(&unexploredTree, rootNode, lastNode)

	assert.NotEmpty(t, path, "Should find path")
	assert.Len(t, path.Nodes, 3, "Should have path with three nodes")
	if 3 == len(path.Nodes) {
		assert.Equal(t, rootNode, path.Nodes[0], "Should have path starting with root node")
		assert.Equal(t, secondNode, path.Nodes[1], "Should have path continuing with second node")
		assert.Equal(t, lastNode, path.Nodes[2], "Should have path finishing with last node")
	}
}

func TestWalkOnCyclingTree(t *testing.T) {
	mockController := gomock.NewController(t)

	neighborFinderMock := generatedMocks.NewMockINeighborFinder(mockController)

	walker := NewBfsWalker(neighborFinderMock)

	unexploredTree := tree.Tree{
		Nodes: []tree.Node{
			tree.Node{ID: 0},
			tree.Node{ID: 1},
			tree.Node{ID: 2},
			tree.Node{ID: 3},
			tree.Node{ID: 4},
			tree.Node{ID: 5},
		},
	}

	rootNode := &unexploredTree.Nodes[0]
	lastNode := &unexploredTree.Nodes[5]

	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, rootNode).
		Return([]*tree.Node{&unexploredTree.Nodes[1], &unexploredTree.Nodes[2], &unexploredTree.Nodes[3]})
	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, &unexploredTree.Nodes[1]).
		Return([]*tree.Node{rootNode, &unexploredTree.Nodes[2], &unexploredTree.Nodes[3]}).
		MaxTimes(3)
	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, &unexploredTree.Nodes[2]).
		Return([]*tree.Node{rootNode, &unexploredTree.Nodes[1], &unexploredTree.Nodes[4]}).
		MaxTimes(3)
	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, &unexploredTree.Nodes[3]).
		Return([]*tree.Node{rootNode, &unexploredTree.Nodes[1], &unexploredTree.Nodes[4]}).
		MaxTimes(3)
	neighborFinderMock.
		EXPECT().
		GetNeighbors(&unexploredTree, &unexploredTree.Nodes[4]).
		Return([]*tree.Node{rootNode, &unexploredTree.Nodes[2], &unexploredTree.Nodes[3], &unexploredTree.Nodes[5]}).
		MaxTimes(2)

	path := walker.Walk(&unexploredTree, rootNode, lastNode)

	assert.NotEmpty(t, path, "Should find path")
	assert.Len(t, path.Nodes, 4, "Should have path with for nodes")
	if 4 == len(path.Nodes) {
		assert.Equal(t, rootNode, path.Nodes[0], "Should have path starting with root node")
		assert.NotEqual(t, &unexploredTree.Nodes[1], path.Nodes[1], "Should have path not using first node")
	}
}
