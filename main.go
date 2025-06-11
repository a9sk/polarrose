// polarrose -- v0.1.0 -- polar rose pattern generator
//
//   Command line system information tool using mathematical
//   rose patterns generation in ASCII art.
//   Copyright (c) 2025, Emiliano Rizzonelli
//
// More info: https://en.wikipedia.org/wiki/Rose_(mathematics)

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a9sk/polarrose/cmd"
	"github.com/a9sk/polarrose/internal/models"
)

// Options:
// -color       color of the rose (ANSI color name)
// -size        radius of the rose (default 20)
// -petals      number of petals (default 5)
// TODO: add more specific and cool flags for a nicer generation
var (
	size   = flag.Int("size", 20, "Size of the rose (radius)")
	color  = flag.String("color", "blue", "Color of the rose (ANSI color name)")
	petals = flag.Int("petals", 5, "Number of petals (default 5)")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: rose [options]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n\n")
		flag.PrintDefaults()
	}

	// parse the flags before doing anything else
	flag.Parse()

	if *size <= 0 {
		fmt.Printf("Error: Size must be a positive integer.\n")
		return
	}

	if *petals <= 0 {
		fmt.Printf("Error: Number of petals must be a positive integer.\n")
		return
	}

	if !models.Colors[*color] {
		fmt.Fprintf(os.Stderr, "Error: the selected \"%s\" color is not supported.\n", *color)
	}

	cmd.Root()
}
