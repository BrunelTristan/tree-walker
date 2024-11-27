package walker

import (
	"tree-walker/model/tree"
)

type IWalker interface {
	Walk(tree tree.Tree)
}
