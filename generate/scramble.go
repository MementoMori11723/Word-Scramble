package generate

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
)

type Word struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Meaning  string `json:"meaning"`
}

func Scramble() Word {
	word, err := get_word()
	if err != nil {
		slog.Error(err.Error())
		fmt.Println("Internal error!\nCan not get a random word!")
		os.Exit(1)
	}

	meaning, err := get_meaning(word)
	if err != nil {
		slog.Error(err.Error())
		meaning = "Can't find the meaning!"
	}

	new_word := []rune(word)
	rand.Shuffle(len(new_word), func(i, j int) {
		new_word[i], new_word[j] = new_word[j], new_word[i]
	})

	return Word{
		Question: string(new_word),
		Answer:   word,
		Meaning:  meaning,
	}
}
