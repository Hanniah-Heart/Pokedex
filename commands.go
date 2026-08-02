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
                }, "mapb": {
                        name: "mapb",
                        description: "Show previous map page",
                        callback: commandMapBack,
                },
        }
        return commandList
}

func commandMap(map_conf *config) error {
	if map_conf.nxt_pg == "" {
		println("There are no further pages")
	} else {
		mapOut(map_conf.nxt_pg, map_conf)
	}
	return nil
}

func commandMapBack(map_conf *config) error {
	if map_conf.prv_pg == "" {
		println("you're on the first page")
	} else {
		mapOut(map_conf.prv_pg, map_conf)
	}
	return nil
}


func mapOut(url string, map_conf *config) error {
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
        var unmarshalledBody mappage
        err2 := json.Unmarshal(body, &unmarshalledBody)
        if err2 != nil {
                fmt.Printf("%v", err2)
                return err
        }
        for i := range unmarshalledBody.Results {
                fmt.Printf("%s\n", unmarshalledBody.Results[i].Name)
        }
	map_conf.nxt_pg = unmarshalledBody.Next
	map_conf.prv_pg = unmarshalledBody.Previous
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
