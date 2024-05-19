package cli

import (
	"fmt"

	"golang.org/x/term"
)

func Section(title string) {
	var dashes []byte
	for range len(title) {
		dashes = append(dashes, '-')
	}

	fmt.Println()
	fmt.Println(title)
	fmt.Println(string(dashes))
}

func Line(text string) {
	fmt.Println(text)
}

func Text(text string) {
	fmt.Print(text)
}

func Debug(text string) {
	fmt.Println("[#] " + text)
}

func Error(text string) {
	fmt.Println("\033[31m[✕] " + text + "\033[m")
}

func Success(text string) {
	fmt.Println("\033[32m[✓] " + text + "\033[m")
}

func Notice(text string) {
	fmt.Println("[\033[32m+\033[m] " + text)
}

func Warning(text string) {
	fmt.Println("\033[33m[!] " + text + "\033[m")
}

func Note(text string) {
	fmt.Println("\033[34m[~] " + text + "\033[m")
}

func Ask(question string) string {
	var input string

	for {
		fmt.Print("[?] " + question)
		fmt.Scan(&input)

		if input == "" {
			Error("The value cannot be empty.")
			continue
		}

		return input
	}
}

func AskHidden(question string) string {
	for {
		fmt.Print("[¿] " + question)
		input, err := term.ReadPassword(0)
		fmt.Print("\n")

		if err != nil {
			Error("Error reading the value.")
			continue
		}

		if string(input) == "" {
			Error("The value cannot be empty.")
			continue
		}

		return string(input)
	}
}
