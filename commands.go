package main
import (
	"fmt"
	"os"
	"io"
	"log"
	"net/http"
	"encoding/json"
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
	url := "https://pokeapi.co/api/v2/location-area/"
	mapout(url)
	return nil
}

func mapout(url string) error {
	res, err := http.Get(url)
	if err != nil { //check for errors with the url
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 { //Check for error related status codes
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil { //check for other errors with reading
		log.Fatal(err)
	}
	type result struct {
                Name    string
                Url     string
        }
        type mappage struct {
                Count           int
                Next            string
                Previous        string
                Results         []result
        }
        var unmarshalledBody mappage
        err2 := json.Unmarshal(body, &unmarshalledBody)
        if err2 != nil {
                fmt.Printf("%v", err2)
                return err
        }
        for i := range unmarshalledBody.Results {
                fmt.Printf("%s\n", unmarshalledBody.Results[i].Name)
        }
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
