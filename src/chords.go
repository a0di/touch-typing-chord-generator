package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
  "math/rand"
)

// Type for layout where fingers are numbers instead of strings -> array of string arrays
type Layout [][]string

// Input parameters
type params struct {
	layoutSelection string
	minChordLen     int
	maxChordLen     int
	repeatChords    int
	shuffle         string
}

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

// ideas for more checks:
// - repeated keys as separate chords
// - calculate distance between keys

func contains(slice []string, key string) bool {
	// Check if the key is in the slice
	for _, value := range slice {
		if value == key {
			return true
		}
	}
	return false
}

func checkIfChordEnds(currentKey string, nextKey string, layout Layout) bool {
	// Check if the current key and next key are assigned to the same finger byt checking layout
	// If they are, the chord ends
	result := false

	// Convert the keys to uppercase, take last character
	currentKey = strings.ToUpper(currentKey)
	currentKey = string(currentKey[len(currentKey)-1])
	nextKey = strings.ToUpper(nextKey)

	for _, finger := range layout {
		if contains(finger, currentKey) && contains(finger, nextKey) {
			result = true
			break
		}
	}
	return result
}

func transformLayoutFromMap(layout map[string][]string) Layout {
	// Transform the layout from map to Layout type
	result := Layout{}
	// Iterate over the map and append the values to the result
	for _, value := range layout {
		result = append(result, value)
	}
	return result
}

func splitWordToChords(layout map[string][]string, word string, params params) []string {
	// A chord is a combination of keys that are pressed at the same time
	// Chord cannot be two keys that are assigned to the same finger
	chords := []string{}
	transformedLayout := transformLayoutFromMap(layout)

	currentKey := string(word[0])

	for i := 1; i < len(word); i++ {
		nextKey := string(word[i])
		// Check if the chord ends
		if checkIfChordEnds(currentKey, nextKey, transformedLayout) || len(currentKey) >= params.maxChordLen {
			// If it does, append the chord to the result and start a new one
			if len(currentKey) >= params.minChordLen {
				for j := 0; j < params.repeatChords; j++ {
					chords = append(chords, currentKey)
				}
			}
			currentKey = nextKey
		} else {
			// If it doesn't, add the key to the current chord
			currentKey += nextKey
		}
	}
	if len(currentKey) >= params.minChordLen {
		for j := 0; j < params.repeatChords; j++ {
			chords = append(chords, currentKey)
		}
	}

	return chords
}

func getParams() params {
	// Get the parameters from the user
	var layoutSelection string
	var minChordLength int
	var maxChordLength int
	var repeatChords int
	var shuffle string

	fmt.Println("Which layout would you like to use?")
	fmt.Println("0 - QWERTY")
	fmt.Println("1 - Dvorak")
	fmt.Println("x - Custom")
	fmt.Scanln(&layoutSelection)

	fmt.Println("Enter the minimum chord length: (default 2)")
	fmt.Scanln(&minChordLength)

	fmt.Println("Enter the maximum chord length: (default 4)")
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
		maxChordLength = 4
	}
	if repeatChords == 0 {
		repeatChords = 1
	}
	if len(shuffle) == 0 {
		shuffle = "n"
	}

	return params{layoutSelection, minChordLength, maxChordLength, repeatChords, shuffle}
}

func main() {
	var layout map[string][]string
	var result []string

	// Get the parameters from the user
	params := getParams()

	// If custom, ask to enter the keys row by row, every row ends with space
	if params.layoutSelection == "x" {
		// Custom layout should be in a file e.g. custom.txt
		// Ask the user to enter the file name
		fmt.Println("Enter the file name for the custom layout: ")
		var custom_layout string
		fmt.Scanln(&custom_layout)
	}

	// If qwerty, get the qwerty layout from file
	if params.layoutSelection == "0" {
		layout = getLayoutAsMap("layouts/qwerty.json")
	}

	// Get the chords based on the layout and file list
	wordlist := getWords()

	// Shuffle array if needed
	if params.shuffle == "y" {
		rand.Shuffle(len(wordlist), func(i, j int) {
			wordlist[i], wordlist[j] = wordlist[j], wordlist[i]
		})
	}

	// Iterate over the word list and get the chords
	for _, word := range wordlist {
		chords := splitWordToChords(layout, word, params)
		result = append(result, chords...)
	}
	fmt.Println(result)
}
