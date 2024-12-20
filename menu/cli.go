package menu

import (
	"Word-scramble/generate"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Cli() {
	fmt.Println("Welcome to Word Scramble Game!")
	fmt.Println("Loading game...")
	word := generate.Scramble()
	fmt.Println("Game loaded!")
	time.Sleep(5 * time.Millisecond)
	fmt.Println("\033c")
	fmt.Println("===================================")
	fmt.Println("         WORD SCRAMBLE GAME        ")
	fmt.Println("===================================")
	fmt.Printf("Scrambled Word: %s\n", word.Question)
	fmt.Println("-----------------------------------")
	reader := bufio.NewReader(os.Stdin)
	for attempt := 1; attempt <= 3; attempt++ {
		fmt.Printf("Attempt %d of 3\n", attempt)
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		answer := strings.TrimSpace(input)
		if strings.EqualFold(answer, word.Answer) {
			fmt.Println("\nCorrect! Well Done!")
			fmt.Printf("Meaning: %s\n", word.Meaning)
			fmt.Println("===================================")
			return
		}
		fmt.Println("Incorrect!")
		fmt.Printf("Attempts Left: %d\n", 3-attempt)
		fmt.Println("-----------------------------------")
	}
	fmt.Println("\nGame Over!")
	fmt.Printf("Correct Answer: %s\n", word.Answer)
	fmt.Printf("Meaning: %s\n", word.Meaning)
	fmt.Println("===================================")
}
