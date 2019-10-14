package main

import (
	"Hangman/dictionnary"
	"Hangman/hangman"
	"fmt"
	"os"
)

func main() {

	err := dictionnary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load the dictionnary: %v\n", err)
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionnary.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
