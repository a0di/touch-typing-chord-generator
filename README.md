# Touch typing chord generator
Generate chords for touch typing practice for QWERTY or alternative layouts based on 1000 most common words or custom wordlists.

> A chord is a combination of keys that are pressed almost simultaneously when typing.

Chords are essential to reach high typing speeds of 60 wpm and above.
This tool allows to generate chords of any lengths (2, 3, 4 keys etc.).

## Installation
```
go build ./src/*
./chord_picker --help
```

## Usage
- generate the list of chords using this tool
- input the result to the tool like [monkeytype](https://monkeytype.com/) or [keybr](https://www.keybr.com/)
- practice, practice, practice!

## Customization
### Input text
By default it uses the list of 1k most common words, but you can provide your own list as an input.

### Repeated chords
Chords in the output can be repeated as many times as you specify, to aid learning and muscle memory.

### Blacklist
Blacklist contains combos that are too awkward and if the chord contains any blacklisted combo it will be omitted.

### Chord lengths
You can define a range of desired chord lengths.

## Alternative layouts

## Adding your custom layout

