package chord_picker

import (
	"strings"
)

// Check if the slice contains the keys
func contains(slice []string, keys string) bool {
	// Split string into characters
	k := strings.Split(keys, "")
	// Check if the key is in the slice
	for _, value := range slice {
		for _, key := range k {
			if value == key {
				return true
			}
		}
	}
	return false
}

// Check if the chord ends based on the layout and the keys
func checkIfChordEnds(currentKey string, nextKey string, layout Layout) bool {
	result := false

	currentKey = strings.ToUpper(currentKey)
	nextKey = strings.ToUpper(nextKey)

	for _, finger := range layout {
		if contains(finger, currentKey) && contains(finger, nextKey) {
			result = true
			break
		}
	}
	return result
}

// Transform the layout from map to Layout type
func transformLayoutFromMap(layout map[string][]string) Layout {
	result := Layout{}
	// Iterate over the map and append the values to the result
	for _, value := range layout {
		result = append(result, value)
	}
	return result
}

// Check if the chord is in the blacklist
func isChordInBlacklist(chord string, blacklist []string) bool {
	for _, item := range blacklist {
		if strings.Contains(strings.ToUpper(chord), strings.ToUpper(item)) {
			return true
		}
	}
	return false
}

// Split the word into chords based on the layout
func SplitWordToChords(layout map[string][]string, word string, params params, blacklist []string) []string {
	// A chord is a combination of keys that are pressed at the same time
	// Chord cannot contain two keys that are assigned to the same finger
	chords := []string{}
	transformedLayout := transformLayoutFromMap(layout)

	currentKey := string(word[0])

	for i := 1; i < len(word); i++ {
		nextKey := string(word[i])

		if checkIfChordEnds(currentKey, nextKey, transformedLayout) || len(currentKey) >= params.maxChordLen {
			if len(currentKey) >= params.minChordLen && !isChordInBlacklist(currentKey, blacklist) {
				for j := 0; j < params.repeatChords; j++ {
					chords = append(chords, currentKey)
				}
			}
			currentKey = nextKey
		} else {
			currentKey += nextKey
		}
	}

	if len(currentKey) >= params.minChordLen && !isChordInBlacklist(currentKey, blacklist) {
		for j := 0; j < params.repeatChords; j++ {
			chords = append(chords, currentKey)
		}
	}

	return chords
}
