package main

import (
    "fmt"
    "os"
    "bufio"
    "encoding/json"
)

/* layout looks like this
{
  "thumbLeft": ["Space", "Alt", "Command", "Option"],
  "thumbRight": ["Space"],
  "indexFingerLeft": ["R", "T", "F", "G", "V", "B", "4", "5"],
  "indexFingerRight": ["Y", "U", "H", "J", "N", "M", "6", "7"],
  "middleFingerLeft": ["E", "D", "C", "3"],
  "middleFingerRight": ["I", "K", ",", "8"],
  "ringFingerLeft": ["W", "S", "X", "2"],
  "ringFingerRight": ["O", "L", ".", "9"],
  "pinkyFingerLeft": ["Q", "A", "Z", "1", "Shift", "Tab", "CapsLock"],
  "pinkyFingerRight": ["P", ";", "/", "-", "=", "[", "]", "\\", "Shift", "Enter", "Backspace"]
}
*/

// Get the layout from json file based on layout above and store as a map
func getLayout(file string) (layout map[string][]string) {
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

// Get the word list from file
func getWords() (wordlist []string) {
  // Open the file
  file, err := os.Open("wordlists/common-words100.txt")
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


func getAdjacentKeys(layout[] string, key string) {

}


func splitWordToChords(layout[] string, word string) {
  // A chord is a combination of keys that are pressed at the same time
  // Chord cannot be two keys that are assigned to the same finger
}


func main() {
  var layoutSelection string

  // Prompt the user which layout to use, 0 - qwerty, 1 - dvorak, x -custom
  fmt.Println("Which layout would you like to use?")
  fmt.Println("0 - QWERTY")
  fmt.Println("1 - Dvorak")
  fmt.Println("x - Custom")

  fmt.Scanln(&layoutSelection)
  fmt.Println("You selected: ", layoutSelection)

  // If custom, ask to enter the keys row by row, every row ends with space
  if layoutSelection == "x" {
    // Custom layout should be in a file e.g. custom.txt
    // Ask the user to enter the file name
    fmt.Println("Enter the file name for the custom layout: ")
    var custom_layout string
    fmt.Scanln(&custom_layout)
  }

  // If qwerty, get the qwerty layout from file
  if layoutSelection == "0" {
    layout := getLayout("layouts/qwerty.json")
    fmt.Println("QWERTY layout selected")
    fmt.Println(layout)
  }

  // Get the chords based on the layout and file list
}
