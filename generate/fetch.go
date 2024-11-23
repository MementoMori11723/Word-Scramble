package generate

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var url = "https://random-word-api.herokuapp.com/word?length=7&lang=en"

func get_word() (string, error) {
  client := &http.Client{
    Timeout: time.Second * 30,
  }
  resp, err := client.Get(url)
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
