package walker

import (
	"tree-walker/model/tree"
)

type BfsWalker struct {
}

func (w BfsWalker) Walk(unexploredTree tree.Tree, start tree.Node, target tree.Node) tree.Path {
	return tree.Path{}
}
