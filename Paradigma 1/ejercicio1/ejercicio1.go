package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ingrese el texto:")
	text := readMultilineInput()

	charCount := len(text)
	wordCount := countWords(text)
	lineCount := countLines(text)

	fmt.Printf("Número de caracteres: %d\n", charCount)
	fmt.Printf("Número de palabras: %d\n", wordCount)
	fmt.Printf("Número de líneas: %d\n", lineCount)
}

func readMultilineInput() string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countLines(text string) int {
	lines := strings.Split(text, "\n")
	return len(lines)
}
