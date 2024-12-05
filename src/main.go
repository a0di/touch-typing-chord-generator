package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"src/chord_picker"
)

// Type for layout where fingers are numbered starting from left pinky
type Layout [][]string

// Input parameters
type params struct {
	layoutSelection string
	minChordLen     int
	maxChordLen     int
	repeatChords    int
	shuffle         string
}

// Read the layout from a json file to a map
func getLayoutAsMap(file string) (layout map[string][]string) {
	// Open the file
	layout = make(map[string][]string)
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer jsonFile.Close()

	// Read the file
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&layout)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	return
}

// Read the word list from a file
func getWords() (wordlist []string) {
	// Open the file
	file, err := os.Open("wordlists/common1000.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
	return
}

// Read the blacklist from a file
func getBlacklist(filename string) (wordlist []string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
	return
}

// Get the parameters from the user
func getParams() params {

	var layoutSelection string
	var minChordLength int
	var maxChordLength int
	var repeatChords int
	var shuffle string

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
	fmt.Println("x - Custom")
	fmt.Scanln(&layoutSelection)

	fmt.Println("Enter the minimum chord length: (default 2)")
	fmt.Scanln(&minChordLength)

	fmt.Println("Enter the maximum chord length: (default 5)")
	fmt.Scanln(&maxChordLength)

	fmt.Println("Enter how many times to repeat the chords (default 1)")
	fmt.Scanln(&repeatChords)

	fmt.Println("Shuffle output? [y/n] (default n)")
	fmt.Scanln(&shuffle)

	// Set default vals if no input
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
	if len(shuffle) == 0 {
		shuffle = "n"
	}

	return params{layoutSelection, minChordLength, maxChordLength, repeatChords, shuffle}
}

// Output the found chords to a file
func outputToFile(filename string, values []string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value)
	}
	return nil
}

func main() {
	var layoutJSON string
	var layout map[string][]string
	var result []string
	outFileName := "output.txt"

	// Get the parameters from the user
	params := getParams()

	// Select layout
	switch params.layoutSelection {
	case "0":
		layoutJSON = "layouts/qwerty.json"
	case "1":
		layoutJSON = "layouts/workman.json"
	case "2":
		layoutJSON = "layouts/dvorak.json"
	case "3":
		layoutJSON = "layouts/programmer-dvorak.json"
	case "4":
		layoutJSON = "layouts/colemak.json"
	case "5":
		layoutJSON = "layouts/colemak-mod-dh.json"
	case "6":
		layoutJSON = "layouts/qwertz.json"
	case "7":
		layoutJSON = "layouts/halmak.json"
	case "8":
		layoutJSON = "layouts/engram.json"
	case "9":
		layoutJSON = "layouts/maltron.json"
	case "10":
		layoutJSON = "layouts/norman.json"
	case "11":
		layoutJSON = "layouts/qgmlw.json"
	case "x":
		// Custom layout should be in a file e.g. custom.txt
		// Ask the user to enter the file name
		fmt.Println("Enter the file name for the custom layout: ")
		fmt.Scanln(&layoutJSON)
	default:
		fmt.Println("Unknown layout")
		return
	}

	layout = getLayoutAsMap(layoutJSON)
	if layout == nil {
		fmt.Println("Layout not found")
		return
	}

	// Get the chords based on the layout and file list
	wordlist := getWords()

	// Get the blacklist based on basename of path to layout
	blacklist := getBlacklist("blacklists/" + layoutJSON[8:len(layoutJSON)-5] + ".txt")

	// Shuffle array if needed
	if params.shuffle == "y" {
		rand.Shuffle(len(wordlist), func(i, j int) {
			wordlist[i], wordlist[j] = wordlist[j], wordlist[i]
		})
	}

	// Iterate over the word list and get the chords
	for _, word := range wordlist {
		chords := chord_picker.SplitWordToChords(layout, word, params, blacklist)
		result = append(result, chords...)
	}

	outputToFile(outFileName, result)
	fmt.Println("Number of chords found: ", len(result)/params.repeatChords)
	fmt.Println("Output saved to: ", outFileName)
}