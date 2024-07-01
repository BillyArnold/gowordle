package main

import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
  	"time"
	"github.com/fatih/color"
	"strings"
)

func main () {
	lineLetters, guessTotal := setRules()
	fmt.Printf("Playing Wordle:\n%d letters \n%d guesses\n", lineLetters, guessTotal)
	possibleWords := getWordsByLength(lineLetters);

	//get random word
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	selectedWord := possibleWords[rand.Intn(len(possibleWords))]

	gameResult := false

	for i := 1; i < guessTotal + 1; i++ {
		attemptResult := guessAttempt(selectedWord)

		if (attemptResult == true) {
			c := color.New(color.FgGreen).Add(color.Underline)
			c.Println("\nCorrect!")
			gameResult = true
			break
		}
		
	}

	if (gameResult == true) {
		c := color.New(color.FgGreen).Add(color.Underline)
		c.Println("\nYou Win!")
		return
	}
		
	c := color.New(color.FgRed).Add(color.Underline)
	c.Println("\nOut of guesses, You lose!")
	c.Printf("The word was : %s\n", selectedWord)
	return
	
}

func setRules () (int, int) {
	c := color.New(color.FgGreen).Add(color.Underline)
	c.Println("Welcome to wordle cli")

	lettersInput := intInput("Select a number of letters: ", 5)
	guessesInput := intInput("Select a number of guesses to allow: ", 5);

	return lettersInput, guessesInput
}

func guessAttempt (wordToGuess string) bool {
	var guess string
	fmt.Println("\nEnter a guess.")
	fmt.Scanln(&guess)

	finalWordLetters := strings.Split(wordToGuess, "")
	guessLetters := strings.Split(guess, "")

	//if guess is wrong length, immediately pass
	if (len(finalWordLetters) != len(guessLetters)) {
		fmt.Printf("Wrong number of letters!\n")
		for i := 1; i < len(finalWordLetters) + 1; i++ {
			fmt.Printf("- ")
		}
		fmt.Printf("\n")

		return false;
	}

	correctLetters := 0

	for i := 0; i < len(finalWordLetters); i++ {
		if (finalWordLetters[i] == guessLetters[i]) {
			c := color.New(color.FgGreen).Add(color.Underline)
			c.Print(guessLetters[i])
			correctLetters++
		} else if (Contains(finalWordLetters, guessLetters[i])) {
			c := color.New(color.FgYellow).Add(color.Underline)
			c.Print(guessLetters[i])
		} else {
			c := color.New(color.FgRed).Add(color.Underline)
			c.Print(guessLetters[i])
		}
	}
	fmt.Printf("\n");
	//if all letters are correct, return true, else return false;
	if(correctLetters == len(finalWordLetters)) {
		return true;
	}
	return false
}

func printResultGrid (results []string) {

}

func intInput (inputText string, inputDefault int) int {
	fmt.Printf("%s", inputText)

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Printf("An error occured while reading input. Number will default to %d)", inputDefault)
		i = 5
	}
	return i
}

func getWordsByLength(wordLength int) []string {
	var wordsWithLength []string
	filePath := "words.txt"
    readFile, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

    for _, line := range fileLines {
		if (line == "") {
			continue
		}

		if (len(line) == wordLength) {
			wordsWithLength = append(wordsWithLength, line)
		}
    }

	return wordsWithLength
}

func Contains(a []string, x string) bool {
	for _, n := range a {
			if x == n {
					return true
			}
	}
	return false
}