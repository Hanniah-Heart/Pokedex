package main
import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	new_text := strings.Fields(text)
	return new_text
}
