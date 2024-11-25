package composition

import (
	"tree-walker/internal/walker"
)

type CompositionRoot struct {
}

func (c *CompositionRoot) Build() {

}

func (c *CompositionRoot) ComposeWalker() *walker.IWalker {
	return new(walker.IWalker)
}
