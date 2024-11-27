package walker

import (
	"tree-walker/model/tree"
)

type IWalker interface {
	Walk(unexploredTree *tree.Tree, start *tree.Node, target *tree.Node) tree.Path
}
