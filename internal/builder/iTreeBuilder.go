package builder

import (
	"tree-walker/model/tree"
)

type ITreeBuilder interface {
	BuildTree() *tree.Tree
}
