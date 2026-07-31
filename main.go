package main
import (
	"fmt"
	"bufio"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}
