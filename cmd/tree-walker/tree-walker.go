package main

import (
	"fmt"
	"tree-walker/internal/composition"
)

func main() {
	fmt.Println("WALK")

	root := composition.NewCompositionRoot()
	root.Build()
}
