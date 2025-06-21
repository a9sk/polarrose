package models

import "fmt"

// list of supported colors. TODO: add more colors to this list, maybe move it somewhere better
var Colors = map[string]bool{
	"black":   true,
	"red":     true,
	"green":   true,
	"yellow":  true,
	"blue":    true,
	"magenta": true,
	"cyan":    true,
	"white":   true,
}

// map color names to ANSI escape codes
var ColorCodes = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"reset":   "\033[0m",
}

var CurrentColor = "white"

func SetColor(c string) error {

	if Colors[c] {
		CurrentColor = c
		return nil
	}

	return fmt.Errorf("invalid color: %s", c)
}
