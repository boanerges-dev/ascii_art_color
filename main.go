package main

import (
	"fmt"
	"os"
	"strings"
)

// ANSI color codes
var colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"orange":  "\033[38;5;214m",
	"reset":   "\033[0m",
}
// Usage instructions (must match audit EXACTLY)
func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
}

func main() {
	args := os.Args[1:] // skip program name

	var color = "red"   // default color
	var substring = ""  // substring to color
	var inputText string

	// CASE 1: only text → color everything with default red
	if len(args) == 1 {
		inputText = args[0]

	// CASE 2: --color=red "text" → color everything
	} else if len(args) == 2 && strings.HasPrefix(args[0], "--color=") {
		color = strings.TrimPrefix(args[0], "--color=")
		inputText = args[1]

	// CASE 3: substring + text → default red
	} else if len(args) == 2 {
		substring = args[0]
		inputText = args[1]

	// CASE 4: --color=red substring text → color only substring
	} else if len(args) == 3 && strings.HasPrefix(args[0], "--color=") {
		color = strings.TrimPrefix(args[0], "--color=")
		substring = args[1]
		inputText = args[2]

	// ❌ INVALID input → print usage (audit requirement)
	} else {
		printUsage()
		return
	}

	// Generate ASCII art
	result := printASCIIArt(inputText, substring, color)
	fmt.Print(result)
}

// Print ASCII-art with optional coloring
func printASCIIArt(text, substring, color string) string {

	// Read ASCII font file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading standard.txt")
		return ""
	}

	lines := strings.Split(string(data), "\n")

	// Get color code (fallback to red if invalid color)
	colorCode, ok := colors[color]
	if !ok {
		colorCode = colors["red"]
	}

	reset := colors["reset"]

	result := strings.Builder{}

	// Handle newline input (\n)
	rows := strings.Split(text, "\\n")

	for _, word := range rows {

		// Each ASCII char is 8 lines tall
		for row := 0; row < 8; row++ {
			i := 0

			for i < len(word) {

				// ✅ If substring exists and matches → color ONLY substring
				if substring != "" &&
					i+len(substring) <= len(word) &&
					word[i:i+len(substring)] == substring {

					for j := 0; j < len(substring); j++ {
						char := word[i+j]
						index := (int(char)-32)*9 + row + 1

						if index < len(lines) {
							result.WriteString(colorCode + lines[index] + reset)
						}
					}

					i += len(substring)

				// ✅ If NO substring → color EVERYTHING
				} else if substring == "" {

					char := word[i]
					index := (int(char)-32)*9 + row + 1

					if index < len(lines) {
						result.WriteString(colorCode + lines[index] + reset)
					}
					i++

				// ✅ Normal (no color applied)
				} else {
					char := word[i]
					index := (int(char)-32)*9 + row + 1

					if index < len(lines) {
						result.WriteString(lines[index])
					}
					i++
				}
			}

			result.WriteString("\n")
		}
	}

	return result.String()
}