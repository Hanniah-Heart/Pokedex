package main
import (
	"fmt"
	"os"
)

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

func commandExit(map_conf *config) error {
        fmt.Printf("%v", *map_conf)
        fmt.Println("Closing the Pokedex... Goodbye!")
        os.Exit(0)
        return nil
}

func commandHelp(map_conf *config) error {
        fmt.Printf("%v", *map_conf)
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
