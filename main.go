package main
import (
	"fmt"
	"bufio"
	"os"
)

func commandExit(debug_string string) error {
	println(debug_string)
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(debug_string string) error {
	println(debug_string)
	commandList := getCommandList()
	fmt.Println("# Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("## Usage:")
	fmt.Println()
	for _, command := range commandList {
		fmt.Printf("- %v: %v\n", command.name, command.description)
	}
	return nil
}

type config struct {
	prv_pg string
	nxt_pg string
}

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}

func getCommandList() map[string]cliCommand {
	commandList := map[string]cliCommand{
		"help": {
			name: "help",
			description: "Describe how to use the Pokedex",
			callback: commandHelp,
		}, "exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
	return commandList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	debug_string := "debug string 2"
	for true {
		commandList := getCommandList()
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Reached End of File")
			os.Exit(0)
		}
		if len(cleanInput(scanner.Text())) < 1 {
			fmt.Println("Please enter a command")
			continue
		}
		userCommand := (cleanInput(scanner.Text()))[0]
		cmd, exists := commandList[userCommand]
		if exists == false {
			fmt.Print("Unknown command")
		} else {
			err := cmd.callback(debug_string)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}
