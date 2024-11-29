package treeHelpers

import (
	"tree-walker/model/tree"
)

type INeighborFinder interface {
	GetNeighbors(sourceTree *tree.Tree, currentNode *tree.Node) []*tree.Node
}
