package treeHelpers

import (
	"tree-walker/model/tree"
)

type NeighborFinder struct {
}

func NewNeighborFinder() *NeighborFinder {
	return &NeighborFinder{}
}

func (nf NeighborFinder) GetNeighbors(sourceTree *tree.Tree, currentNode *tree.Node) []*tree.Node {
	if sourceTree == nil || currentNode == nil || len(sourceTree.Links) == 0 {
		return nil
	}

	neighbors := []*tree.Node{}

	for _, link := range sourceTree.Links {
		if link.Nodes[0] == currentNode {
			neighbors = append(neighbors, link.Nodes[1])
		} else if link.Nodes[1] == currentNode {
			neighbors = append(neighbors, link.Nodes[0])
		}
	}

	return neighbors
}
