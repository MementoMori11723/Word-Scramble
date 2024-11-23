package generate

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	random_word_url = "https://random-word-api.herokuapp.com/word?length=7&lang=en"
	dictionary_url = "https://api.dictionaryapi.dev/api/v2/entries/en/"
)

func get_word() (string, error) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Get(random_word_url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var word []string
	err = json.Unmarshal(respBody, &word)
	if err != nil {
		return "", err
	}
	return word[0], nil
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
