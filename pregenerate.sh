#!/bin/bash
# Pregenerate chord lists

# Option to build as well, to do everyting in one go
BUILD=$0

if [[ $BUILD ]]; then
  go build src/*
fi

# QWERTY
./chord_picker -l 0 -min 2 -max 2 -r 1 -s -d
mv output.txt pregenerated/0-qwerty_bigrams_norepeat_shuffle_space.txt
./chord_picker -l 0 -min 3 -max 3 -r 1 -s -d
mv output.txt pregenerated/0-qwerty_trigrams_norepeat_shuffle_space.txt
./chord_picker -l 0 -min 4 -max 4 -r 1 -s -d
mv output.txt pregenerated/0-qwerty_fourgrams_norepeat_shuffle_space.txt
./chord_picker -l 0 -min 2 -max 5 -r 1 -s -d
mv output.txt pregenerated/0-qwerty_min2_max5_norepeat_shuffle_space.txt

./chord_picker -l 0 -min 2 -max 2 -r 3 -s -d
mv output.txt pregenerated/0-qwerty_bigrams_repeat3_shuffle_space.txt
./chord_picker -l 0 -min 3 -max 3 -r 3 -s -d
mv output.txt pregenerated/0-qwerty_trigrams_repeat3_shuffle_space.txt
./chord_picker -l 0 -min 4 -max 4 -r 3 -s -d
mv output.txt pregenerated/0-qwerty_fourgrams_repeat3_shuffle_space.txt
./chord_picker -l 0 -min 2 -max 5 -r 3 -s -d
mv output.txt pregenerated/0-qwerty_min2_max5_repeat3_shuffle_space.txt

./chord_picker -l 0 -min 2 -max 2 -r 1 -s
mv output.txt pregenerated/0-qwerty_bigrams_norepeat_shuffle_nl.txt
./chord_picker -l 0 -min 3 -max 3 -r 1 -s
mv output.txt pregenerated/0-qwerty_trigrams_norepeat_shuffle_nl.txt
./chord_picker -l 0 -min 4 -max 4 -r 1 -s
mv output.txt pregenerated/0-qwerty_fourgrams_norepeat_shuffle_nl.txt
./chord_picker -l 0 -min 2 -max 5 -r 1 -s
mv output.txt pregenerated/0-qwerty_min2_max5_norepeat_shuffle_nl.txt

./chord_picker -l 0 -min 2 -max 2 -r 3 -s
mv output.txt pregenerated/0-qwerty_bigrams_repeat3_shuffle_nl.txt
./chord_picker -l 0 -min 3 -max 3 -r 3 -s
mv output.txt pregenerated/0-qwerty_trigrams_repeat3_shuffle_nl.txt
./chord_picker -l 0 -min 4 -max 4 -r 3 -s
mv output.txt pregenerated/0-qwerty_fourgrams_repeat3_shuffle_nl.txt
./chord_picker -l 0 -min 2 -max 5 -r 3 -s
mv output.txt pregenerated/0-qwerty_min2_max5_repeat3_shuffle_nl.txt

# QWERTZ
./chord_picker -l 6 -min 2 -max 2 -r 1 -s -d
mv output.txt pregenerated/6-qwertz_bigrams_norepeat_shuffle_space.txt
./chord_picker -l 6 -min 3 -max 3 -r 1 -s -d
mv output.txt pregenerated/6-qwertz_trigrams_norepeat_shuffle_space.txt
./chord_picker -l 6 -min 4 -max 4 -r 1 -s -d
mv output.txt pregenerated/6-qwertz_fourgrams_norepeat_shuffle_space.txt
./chord_picker -l 6 -min 2 -max 5 -r 1 -s -d
mv output.txt pregenerated/6-qwertz_min2_max5_norepeat_shuffle_space.txt

./chord_picker -l 6 -min 2 -max 2 -r 3 -s -d
mv output.txt pregenerated/6-qwertz_bigrams_repeat3_shuffle_space.txt
./chord_picker -l 6 -min 3 -max 3 -r 3 -s -d
mv output.txt pregenerated/6-qwertz_trigrams_repeat3_shuffle_space.txt
./chord_picker -l 6 -min 4 -max 4 -r 3 -s -d
mv output.txt pregenerated/6-qwertz_fourgrams_repeat3_shuffle_space.txt
./chord_picker -l 6 -min 2 -max 5 -r 3 -s -d
mv output.txt pregenerated/6-qwertz_min2_max5_repeat3_shuffle_space.txt

./chord_picker -l 6 -min 2 -max 2 -r 1 -s
mv output.txt pregenerated/6-qwertz_bigrams_norepeat_shuffle_nl.txt
./chord_picker -l 6 -min 3 -max 3 -r 1 -s
mv output.txt pregenerated/6-qwertz_trigrams_norepeat_shuffle_nl.txt
./chord_picker -l 6 -min 4 -max 4 -r 1 -s
mv output.txt pregenerated/6-qwertz_fourgrams_norepeat_shuffle_nl.txt
./chord_picker -l 6 -min 2 -max 5 -r 1 -s
mv output.txt pregenerated/6-qwertz_min2_max5_norepeat_shuffle_nl.txt

./chord_picker -l 6 -min 2 -max 2 -r 3 -s
mv output.txt pregenerated/6-qwertz_bigrams_repeat3_shuffle_nl.txt
./chord_picker -l 6 -min 3 -max 3 -r 3 -s
mv output.txt pregenerated/6-qwertz_trigrams_repeat3_shuffle_nl.txt
./chord_picker -l 6 -min 4 -max 4 -r 3 -s
mv output.txt pregenerated/6-qwertz_fourgrams_repeat3_shuffle_nl.txt
./chord_picker -l 6 -min 2 -max 5 -r 3 -s
mv output.txt pregenerated/6-qwertz_min2_max5_repeat3_shuffle_nl.txt

# Colemak
./chord_picker -l 4 -min 2 -max 2 -r 1 -s -d
mv output.txt pregenerated/4-colemak_bigrams_norepeat_shuffle_space.txt
./chord_picker -l 4 -min 3 -max 3 -r 1 -s -d
mv output.txt pregenerated/4-colemak_trigrams_norepeat_shuffle_space.txt
./chord_picker -l 4 -min 4 -max 4 -r 1 -s -d
mv output.txt pregenerated/4-colemak_fourgrams_norepeat_shuffle_space.txt
./chord_picker -l 4 -min 2 -max 5 -r 1 -s -d
mv output.txt pregenerated/4-colemak_min2_max5_norepeat_shuffle_space.txt

./chord_picker -l 4 -min 2 -max 2 -r 3 -s -d
mv output.txt pregenerated/4-colemak_bigrams_repeat3_shuffle_space.txt
./chord_picker -l 4 -min 3 -max 3 -r 3 -s -d
mv output.txt pregenerated/4-colemak_trigrams_repeat3_shuffle_space.txt
./chord_picker -l 4 -min 4 -max 4 -r 3 -s -d
mv output.txt pregenerated/4-colemak_fourgrams_repeat3_shuffle_space.txt
./chord_picker -l 4 -min 2 -max 5 -r 3 -s -d
mv output.txt pregenerated/4-colemak_min2_max5_repeat3_shuffle_space.txt

./chord_picker -l 4 -min 2 -max 2 -r 1 -s
mv output.txt pregenerated/4-colemak_bigrams_norepeat_shuffle_nl.txt
./chord_picker -l 4 -min 3 -max 3 -r 1 -s
mv output.txt pregenerated/4-colemak_trigrams_norepeat_shuffle_nl.txt
./chord_picker -l 4 -min 4 -max 4 -r 1 -s
mv output.txt pregenerated/4-colemak_fourgrams_norepeat_shuffle_nl.txt
./chord_picker -l 4 -min 2 -max 5 -r 1 -s
mv output.txt pregenerated/4-colemak_min2_max5_norepeat_shuffle_nl.txt

./chord_picker -l 4 -min 2 -max 2 -r 3 -s
mv output.txt pregenerated/4-colemak_bigrams_repeat3_shuffle_nl.txt
./chord_picker -l 4 -min 3 -max 3 -r 3 -s
mv output.txt pregenerated/4-colemak_trigrams_repeat3_shuffle_nl.txt
./chord_picker -l 4 -min 4 -max 4 -r 3 -s
mv output.txt pregenerated/4-colemak_fourgrams_repeat3_shuffle_nl.txt
./chord_picker -l 4 -min 2 -max 5 -r 3 -s
mv output.txt pregenerated/4-colemak_min2_max5_repeat3_shuffle_nl.txt
