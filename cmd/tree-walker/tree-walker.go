package main

import (
	"fmt"
	"tree-walker/internal/composition"
	"tree-walker/model/configuration"
)

func main() {
	fmt.Println("WALK")

	conf := &configuration.LaunchingConfiguration{}

	root := composition.NewCompositionRoot(conf)
	root.Build()
}
