package walker

import (
	"tree-walker/model/tree"
)

type BfsWalker struct {
}

func (w BfsWalker) Walk(unexploredTree *tree.Tree, start *tree.Node, target *tree.Node) tree.Path {
	resultPath := tree.Path{}

	if start == nil || target == nil {
		return resultPath
	}

	if start == target {
		resultPath.Nodes = append(resultPath.Nodes, start)
	}

	return resultPath
}
