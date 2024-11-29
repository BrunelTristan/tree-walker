package walker

import (
	"tree-walker/internal/treeHelpers"
	"tree-walker/model/tree"
)

type BfsWalker struct {
	neighborsFinder treeHelpers.INeighborFinder
}

func NewBfsWalker(neighborsFinder treeHelpers.INeighborFinder) *BfsWalker {
	return &BfsWalker{neighborsFinder: neighborsFinder}
}

func (w BfsWalker) Walk(unexploredTree *tree.Tree, start *tree.Node, target *tree.Node) *tree.Path {
	resultPath := new(tree.Path)

	if start == nil || target == nil {
		return resultPath
	}

	resultPath.Nodes = append(resultPath.Nodes, start)

	if start == target {
		return resultPath
	}

	neighbors := w.neighborsFinder.GetNeighbors(unexploredTree, start)
	for _, neighbor := range neighbors {
		w.neighborsFinder.GetNeighbors(unexploredTree, neighbor)
	}

	return resultPath
}
