package composition

import (
	"tree-walker/internal/walker"
)

type CompositionRoot struct {
}

func (c *CompositionRoot) Build() {

}

func (c *CompositionRoot) composeWalker() walker.IWalker {
	return new(walker.BfsWalker)
}
