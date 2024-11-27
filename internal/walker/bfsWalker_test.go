package walker

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tree-walker/model/tree"
)

func TestWalk(t *testing.T) {
	walker := BfsWalker{}

	tree := tree.Tree{}

	walker.Walk(tree)
}
