package main

import (
	"flag"
	"fmt"
)

// Input parameters
type Params struct {
	layoutSelection string
	layoutFile      string
	minChordLen     int
	maxChordLen     int
	repeatChords    int
	shuffle         bool
	interactive     bool
}

var availableLayouts string = `
0 - qwerty
1 - workman
2 - dvorak
3 - programmer-dvorak
4 - colemak
5 - colemak-mod-dh
6 - qwertz
7 - halmak
8 - engram
9 - maltron
10 - norman
11 - qgmlw
`

// Get parameters from cli
func GetFlags() Params {
	var layoutSelection = flag.String("l", "0", "Select layout to use, available: "+availableLayouts)
	var layoutFile = flag.String("f", "", "Path to the custom layout JSON file")
	var minChordLength = flag.Int("min", 2, "Minimum chord length")
	var maxChordLength = flag.Int("max", 5, "Maximum chord length")
	var repeatChords = flag.Int("r", 1, "Repat chords in the output")
	var shuffle = flag.Bool("s", false, "Shuffle chords in the output (maintaining repeated chords)")
	var interactive = flag.Bool("i", false, "Enter values interactively")

	flag.Parse()

	return Params{*layoutSelection, *layoutFile, *minChordLength, *maxChordLength, *repeatChords, *shuffle, *interactive}
}

// Get the parameters from the user interactively
func GetParamsInteractively() Params {

	var layoutSelection string
	var minChordLength int
	var maxChordLength int
	var repeatChords int
	var shuffleStr string

	fmt.Println("Which layout would you like to use? (default 0 - QWERTY)")
	fmt.Println("0 - qwerty")
	fmt.Println("1 - workman")
	fmt.Println("2 - dvorak")
	fmt.Println("3 - programmer-dvorak")
	fmt.Println("4 - colemak")
	fmt.Println("5 - colemak-mod-dh")
	fmt.Println("6 - qwertz")
	fmt.Println("7 - halmak")
	fmt.Println("8 - engram")
	fmt.Println("9 - maltron")
	fmt.Println("10 - norman")
	fmt.Println("11 - qgmlw")
	fmt.Scanln(&layoutSelection)

	fmt.Println("Enter the minimum chord length: (default 2)")
	fmt.Scanln(&minChordLength)

	fmt.Println("Enter the maximum chord length: (default 5)")
	fmt.Scanln(&maxChordLength)

	fmt.Println("Enter how many times to repeat the chords (default 1)")
	fmt.Scanln(&repeatChords)

	fmt.Println("Shuffle output? [y/n] (default n)")
	fmt.Scanln(&shuffleStr)

	// Set default values if no input from the user
	if len(layoutSelection) == 0 {
		layoutSelection = "0"
	}
	if minChordLength == 0 {
		minChordLength = 2
	}
	if maxChordLength == 0 {
		maxChordLength = 5
	}
	if repeatChords == 0 {
		repeatChords = 1
	}
	if len(shuffleStr) == 0 {
		shuffleStr = "n"
	}

	var shuffle bool
	if shuffleStr == "y" {
		shuffle = true
	} else {
		shuffle = false
	}

	return Params{layoutSelection, "", minChordLength, maxChordLength, repeatChords, shuffle, false}
}
