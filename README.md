# Touch typing chord generator
Generate chords for touch typing practice for QWERTY or alternative layouts based on 1000 most common words or custom wordlists.

> A chord is a combination of keys that are pressed almost simultaneously when typing.

Chords are essential to reach high typing speeds of 60 wpm and above.
This tool allows to generate chords of any lengths (2, 3, 4 keys etc.).

## Quickstart
If you don't want to bother with go, running it and configuring your own lists, then check the `pregenerated` directory where some generated chord lists are placed, together with parameters that were used to generate them.

You can take the chord list and input into a tool like [monkeytype](https://monkeytype.com/) or [keybr](https://www.keybr.com/) for practicing.

## Installation and usage
```
go build ./src/*
./chord_picker --help
```

## Customization
### Input text
By default it uses the list of 1k most common words, but you can provide your own list as an input using parameter `w`.

### Repeated chords
Chords in the output can be repeated as many times as you specify, to aid learning and muscle memory. The default is 1 but setting it to e.g. 3 may be helpful.

### Blacklist
Blacklist contains combos that are too awkward and if the chord contains any of the blacklisted combos it will be omitted.

Specify the blacklist for desired layout by putting a txt file with the same name as the layout into `blacklists` directory. See the `qwerty.txt` for an example.

### Chord lengths
You can define a range of desired chord lengths: 2 for bigrams, 3 for trigrams, etc.

## Alternative layouts
The tool supports various layouts, and it is easy to add your own layout if needed.

Supported layouts:
- qwerty
- workman
- dvorak
- programmer-dvorak
- colemak
- colemak-mod-dh
- qwertz
- halmak
- engram
- maltron
- norman
- qgmlw

## Adding your custom layout
Add a json file with your layout description (follow example of `layouts/qwerty.json`), and specify the path to that file with input parameter `f`.

