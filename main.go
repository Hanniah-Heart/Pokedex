package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	for true {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin) /*
		Designate variable "scanner" to hold operating system's standard input (os.Stdin)
		as read and sorted by the NewScanner function from the bufio import.

		The Go bufio manual says this about the function NewScanner
			"
		func NewScanner
		added in go1.1
		func NewScanner(r io.Reader) *Scanner

		NewScanner returns a new Scanner to read from r.
		The split function defaults to ScanLines.
			"

		Therefore scanner is now a bufio designed Scanner which listen's to os.Stdin.
		*/
		scanner.Scan() /*
		the .Scan() method has the following description on the Go bufio manual.

		func (*Scanner) Scan ¶
		added in go1.1
		func (s *Scanner) Scan() bool

			"
		Scan advances the Scanner to the next token,
		which will then be available through the Scanner.
		Bytes or Scanner.Text method.
		It returns false when there are no more tokens,
		either by reaching the end of the input or an error.
		After Scan returns false,
		the Scanner.Err method will return any error that occurred during scanning,
		except that if it was io.EOF, Scanner.Err will return nil.
		Scan panics if the split function returns too many empty tokens
		without advancing the input.
		This is a common error mode for scanners.
			"

		Apparently these tokens can be entire lines of commandline text
		that are only complete when sent.

		So, effectively, this line should ask the process to wait until it recieves a full
		"token" or, effectively, a command line input as taken from its source
		(the standard input).
		*/
		output := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s", output[0])
		fmt.Println()
	}
}
