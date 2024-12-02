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

	builtTree := &tree.Tree{
		Nodes: make([]tree.Node, b.nodeCount),
		Links: make([]tree.Link, b.nodeCount-1),
	}
	currentParentIndex := 0

	for node := 0; node < b.nodeCount; node++ {
		builtTree.Nodes[node].ID = node

		if node != 0 {
			builtTree.Links[node-1].Nodes = [2]*tree.Node{&builtTree.Nodes[currentParentIndex], &builtTree.Nodes[node]}

			if node > 1 && builtTree.Links[node-2].Nodes[0] == &builtTree.Nodes[currentParentIndex] {
				currentParentIndex++
			}
		}
	}

	return builtTree
}
