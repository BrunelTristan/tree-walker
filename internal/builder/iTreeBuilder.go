package builder

import (
	"tree-walker/model/tree"
)

type IBuilder interface {
	BuildTree() *tree.Tree
}
