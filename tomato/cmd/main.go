package main

import (
	"fmt"

	"github.com/AAspCodes/caprese/tomato/internal/cut"
	"github.com/AAspCodes/caprese/tomato/pkg/garden"
	"github.com/AAspCodes/caprese/tomato/pkg/kitchen"
)

func main() {
	fmt.Println("hello world")
	tomato := garden.GrowTomato()
	tomatoSlices := cut.Cut(tomato)
	for _, piece := range tomatoSlices {
		fmt.Println(piece)
	}
	kitchen.CleanWorkspace()
}
