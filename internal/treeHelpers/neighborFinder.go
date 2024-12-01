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
	return nil
}
