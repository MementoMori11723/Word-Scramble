package generate

import (
	_ "embed"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
  //go:embed long-words.txt
  wordsFs string
	dictionary_url = "https://api.dictionaryapi.dev/api/v2/entries/en/"
)

func get_word() (string, error) {
  words := strings.Split(strings.TrimSpace(wordsFs), "\n")
  index := rand.Intn(len(words))
	return words[index], nil 
}

func get_meaning(word string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Get(dictionary_url + word)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var meaning []struct {
		Meanings []struct {
			Definitions []struct {
				Definition string `json:"definition"`
			} `json:"definitions"`
		} `json:"meanings"`
	}

	err = json.Unmarshal(respBody, &meaning)
	if err != nil {
		return "", err
	}

	return meaning[0].
		Meanings[0].
		Definitions[0].
		Definition, nil 
}
