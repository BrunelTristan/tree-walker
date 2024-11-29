package walker

import (
	"slices"
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

	alreadySeen := []*tree.Node{}
	toWalk := []*tree.Path{&tree.Path{Nodes: []*tree.Node{start}}}

	for len(toWalk) > 0 {
		currentPath := toWalk[0]
		currentNode := currentPath.Nodes[len(currentPath.Nodes)-1]
		if currentNode == target {
			return currentPath
		}

		alreadySeen = append(alreadySeen, currentNode)
		neighbors := w.neighborsFinder.GetNeighbors(unexploredTree, currentNode)

		for _, node := range neighbors {
			if !(slices.Contains(alreadySeen, node)) {
				toWalk = append(toWalk, &tree.Path{Nodes: append(currentPath.Nodes, node)})
			}
		}

		toWalk = toWalk[1:]
	}

	return resultPath
}
