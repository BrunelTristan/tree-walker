package composition

import (
	"reflect"
	"tree-walker/internal/builder"
	"tree-walker/internal/treeHelpers"
	"tree-walker/internal/walker"
	"tree-walker/model/tree"
)

type CompositionRoot struct {
	singletons map[reflect.Type]interface{}
}

func NewCompositionRoot() *CompositionRoot {
	return &CompositionRoot{
		singletons: make(map[reflect.Type]interface{}),
	}
}

func (c CompositionRoot) Build() {
	c.singletons[reflect.TypeFor[treeHelpers.INeighborFinder]()] = treeHelpers.NewNeighborFinder()
	c.singletons[reflect.TypeFor[builder.ITreeBuilder]()] = builder.NewBinaryTreeBuilder(99)
}

func (c CompositionRoot) ComposeWalker() walker.IWalker {
	return walker.NewBfsWalker(c.singletons[reflect.TypeFor[treeHelpers.INeighborFinder]()].(treeHelpers.INeighborFinder))
}

func (c CompositionRoot) ComposeTree() *tree.Tree {
	return c.singletons[reflect.TypeFor[builder.ITreeBuilder]()].(builder.ITreeBuilder).BuildTree()
}
