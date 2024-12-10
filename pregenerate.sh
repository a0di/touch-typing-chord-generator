#!/bin/bash
# Pregenerate chord lists

# Option to build as well, to do everything in one go
BUILD=$0

if [[ $BUILD ]]; then
  go build src/*
fi

generate_chords() {
  local layout=$1
  local min=$2
  local max=$3
  local repeat=$4
  local shuffle=$5
  local outname=$6
  local space=$7

  ./chord_picker -l $layout -min $min -max $max -r $repeat -s $shuffle $space
  echo "Generated $outname"
  mv output.txt pregenerated/${layout}-${outname}.txt
}

layouts=("0-qwerty" "1-workman" "2-dvorak" "3-programmer-dvorak" "4-colemak" "5-colemak-mod-dh" "6-qwertz" "7-halmak" "8-engram" "9-maltron" "10-norman" "11-qgmlw")
repeats=(1 3 5)
spaces=("-d" "")
space_suffixes=("space" "nl")

for layout in "${layouts[@]}"; do
  layout_num=${layout%%-*}
  layout_name=${layout#*-}

  for repeat in "${repeats[@]}"; do
    for space_index in "${!spaces[@]}"; do
      space=${spaces[$space_index]}
      space_suffix=${space_suffixes[$space_index]}
      echo "Generating $layout_name repeat $repeat shuffle $space_suffix"
      echo "Command line: ./chord_picker -l $layout_num -min 2 -max 4 -r $repeat -s 1 $space"
      echo "Output file: ${layout_name}_${ngram}grams_repeat${repeat}_shuffle_${space_suffix}"

      for ngram in 2 3 4; do
        generate_chords $layout_num $ngram $ngram $repeat -s "${layout_name}_${ngram}grams_repeat${repeat}_shuffle_${space_suffix}" $space
      done
    done
  done
done


