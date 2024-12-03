package main

import (
	"flag"
	"fmt"
	"tree-walker/internal/composition"
	"tree-walker/model/configuration"
)

func main() {
	fmt.Println("WALK")

	conf := &configuration.LaunchingConfiguration{}
	flag.UintVar(&conf.NodeCount, "nodes", 30, "nodes count in the tree")

	flag.Parse()

	root := composition.NewCompositionRoot(conf)
	root.Build()

	walker := root.ComposeWalker()
	tree := root.ComposeTree()

	path := walker.Walk(tree, &tree.Nodes[0], &tree.Nodes[len(tree.Nodes)-1])

	fmt.Printf("Path from Node %d to %d is ", tree.Nodes[0].ID, tree.Nodes[len(tree.Nodes)-1].ID)
	fmt.Println(path)
}
