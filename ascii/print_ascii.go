package ascii

import (
	"fmt"
	"strings"
)

// printWord prints ASCII art for a single word using the provided ASCII map.
func printWord(word string, asciMap [][]string) (string, error) {
	var asciiArt strings.Builder
	for i := 1; i <= 8; i++ {
		for _, char := range word {

			index := int(char - 32)
			if index < 0 || index >= len(asciMap) {
				return "", fmt.Errorf("unknown character: %q", char)
			} else {
				asciiArt.WriteString(asciMap[index][i])
			}
		}
		asciiArt.WriteString("\n")
	}
	return asciiArt.String(), nil
}

// PrintArt prints ASCII art for a given string using the provided ASCII map.
func PrintArt(str string, asciMap [][]string) (string, error) {
	var asciiPrint string
	switch str {
	case "":
		fmt.Print()
	case "\\n":
		fmt.Println()

	case "\\r", "\\f", "\\v", "\\t", "\\b", "\\a":
		return "", fmt.Errorf("error: unsupported escape sequence '%s'", str)
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		s = strings.ReplaceAll(s, "\\r", "\r")
		s = strings.ReplaceAll(s, "\\f", "\f")
		s = strings.ReplaceAll(s, "\\v", "\v")
		s = strings.ReplaceAll(s, "\\t", "\t")
		s = strings.ReplaceAll(s, "\\b", "\b")
		s = strings.ReplaceAll(s, "\\a", "\a")
		words := strings.Split(s, "\n")
		// fmt.Print(words)

		num := 0

		for _, word := range words {
			if word == "" {
				num++
				if num < len(words) {
					fmt.Println()
					continue
				}
			} else {
				asciiAr, err := printWord(word, asciMap)
				if err != nil {
					return "", err
				}
				asciiPrint += asciiAr

			}
		}
	}
	return asciiPrint, nil
}
