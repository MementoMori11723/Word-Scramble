package generate

import (
	"math/rand"
)

func Scramble() (string, string) {
  word, err := get_word()
  if err != nil {
    panic(err)
  }
  new_word := []rune(word)
  rand.Shuffle(len(new_word), func(i, j int) {
    new_word[i], new_word[j] = new_word[j], new_word[i]
  })
  return string(new_word), word
}
