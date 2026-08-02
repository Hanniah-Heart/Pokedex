package main
import (
	"fmt"
	"os"
	"io"
	"log"
	"net/http"
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
                }, "map": {
                        name: "map",
                        description: "Show next map page",
                        callback: commandMap,
                },
        }
        return commandList
}

func commandMap(map_conf *config) error {
	mapout()
	return nil
}

func mapout() error {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
        return nil
}

func commandExit(map_conf *config) error {
        fmt.Println("Closing the Pokedex... Goodbye!")
        os.Exit(0)
        return nil
}

func commandHelp(map_conf *config) error {
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
