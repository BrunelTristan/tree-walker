package main

import (
	"testing"
	"tree-walker/model/configuration"
)

func TestMain(t *testing.T) {
	main()
}

func TestLaunchForVersion(t *testing.T) {
	launch(true, nil)
}

func TestLaunchForWalk(t *testing.T) {
	conf := configuration.LaunchingConfiguration{}

	conf.NodeCount = 6

	launch(false, &conf)
}
