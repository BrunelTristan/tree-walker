package builder

import (
	"tree-walker/model/tree"
)

type BinaryTreeBuilder struct {
	nodeCount int
}

func NewBinaryTreeBuilder(count int) *BinaryTreeBuilder {
	return &BinaryTreeBuilder{
		nodeCount: count,
	}
}

func (b BinaryTreeBuilder) BuildTree() *tree.Tree {
	if b.nodeCount == 0 {
		return nil
	}

	builtTree := new(tree.Tree)
	currentParentIndex := 0
	firstLink := true

	for node := 0; node < b.nodeCount; node++ {
		builtTree.Nodes = append(builtTree.Nodes, tree.Node{ID: node})

		if node != 0 {
			builtTree.Links = append(builtTree.Links, tree.Link{
				Nodes: [2]*tree.Node{&builtTree.Nodes[currentParentIndex], &builtTree.Nodes[len(builtTree.Nodes)-1]},
			})

			firstLink = !firstLink
			if firstLink {
				currentParentIndex++
			}
		}
	}

	return builtTree
}
