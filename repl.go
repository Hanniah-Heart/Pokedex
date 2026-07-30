package main
import "strings"
import "fmt"

func cleanInput(text string) []string {
	fmt.Printf("the original input is %v \n", text)
	text = strings.ToLower(text)
	fmt.Printf("the lowercase text is %v \n", text)
	text = strings.TrimSpace(text)
	fmt.Printf("the trimmed text is %v \n", text)
	new_text := strings.Fields(text)
	fmt.Printf("the Fields of the text are now %v \n", new_text)
	return new_text
}
