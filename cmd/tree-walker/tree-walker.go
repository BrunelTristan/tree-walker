package main

import (
	"flag"
	"fmt"
	"tree-walker/internal/composition"
	"tree-walker/model/configuration"
)

const (
	major = 0
	minor = 1
	patch = 1
)

func readFlags() (shouldDisplayVersion bool, conf *configuration.LaunchingConfiguration) {
	conf = &configuration.LaunchingConfiguration{}
	flag.UintVar(&conf.NodeCount, "nodes", 30, "nodes count in the tree")
	shouldDisplayVersion = *flag.Bool("v", false, "display version")

	flag.Parse()

	return
}

func launch(shouldDisplayVersion bool, conf *configuration.LaunchingConfiguration) {
	if shouldDisplayVersion {
		displayVersion()
	} else {
		launchWalk(conf)
	}
}

func displayVersion() {
	fmt.Printf("V%d.%d.%d\n", major, minor, patch)
}

func launchWalk(conf *configuration.LaunchingConfiguration) {
	root := composition.NewCompositionRoot(conf)
	root.Build()

	walker := root.ComposeWalker()
	tree := root.ComposeTree()

	path := walker.Walk(tree, &tree.Nodes[0], &tree.Nodes[len(tree.Nodes)-1])

	fmt.Printf("Path from Node %d to %d is ", tree.Nodes[0].ID, tree.Nodes[len(tree.Nodes)-1].ID)
	fmt.Println(path)
}

func main() {
	shouldDisplayVersion, conf := readFlags()

	launch(shouldDisplayVersion, conf)
}

// TODO : implement concurrent BFS with go routine to explore each floor concurrently (check benchmark)
// TODO : change walker to walk along all the tree (with a functor)
// TODO : create a path finder
// TODO : immplement DFS
// TODO : manage weighted links
// TODO : implement IsConvex
