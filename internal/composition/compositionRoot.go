package composition

import (
	"reflect"
	"tree-walker/internal/treeHelpers"
	"tree-walker/internal/walker"
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
}

func (c CompositionRoot) composeWalker() walker.IWalker {
	return walker.NewBfsWalker(c.singletons[reflect.TypeFor[treeHelpers.INeighborFinder]()].(treeHelpers.INeighborFinder))
}
