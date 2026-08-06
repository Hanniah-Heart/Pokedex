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
                }, "explore": {
                        name: "explore",
                        description: "Show information about location named in first argument",
                        callback: commandExplore,
                },
        }
        return commandList
}

func commandExplore(map_conf *config, args []string) error {
        if len(args) != 1 {
		fmt.Println("expected 1 argument recieved ", len(args), "\n", args)
		return nil
	}
	var unmarshalledBody locationArea
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	if err := callURL(map_conf, url, &unmarshalledBody); err != nil { return err }
	fmt.Println("Exploring ", args[0], "...")
	fmt.Println("Found Pokemon:")
	output := unmarshalledBody.Pokemon_Encounters
	for i := range output {
		fmt.Printf("%s\n", output[i].Pokemon.Name)
	}
	return nil
}

func callURL(map_conf *config, url string, unmarshalledBody any) (error) {
	if body, ok := map_conf.CacheAddress.Get(url); ok {
		if err := json.Unmarshal(body, &unmarshalledBody); err != nil { return err }
	} else {
		body, err := getFromWeb(url)
		if err != nil {	return err }
		if err = map_conf.CacheAddress.Add(url, body); err != nil { return err }
	        if err = json.Unmarshal(body, &unmarshalledBody); err != nil { return err }
	}
	return nil
}

func commandMap(map_conf *config, args []string) error {
        if len(args) >= 1 {
		fmt.Println("expected no arguments recieved ", len(args), "\n", args)
	}
	if map_conf.nxt_pg == "" {
		println("There are no further pages")
	} else {
		mapOut(map_conf.nxt_pg, map_conf)
	}
	return nil
}

func commandMapBack(map_conf *config, args []string) error {
        if len(args) >= 1 {
		fmt.Println("expected no arguments recieved ", len(args), "\n", args)
	}
	if map_conf.prv_pg == "" {
		println("you're on the first page")
	} else {
		mapOut(map_conf.prv_pg, map_conf)
	}
	return nil
}


func mapOut(url string, map_conf *config) error {
	var unmarshalledBody mappage
	if err := callURL(map_conf, url, &unmarshalledBody); err != nil {return err}
	for i := range unmarshalledBody.Results {
		fmt.Printf("%s\n", unmarshalledBody.Results[i].Name)
	}
	map_conf.nxt_pg = unmarshalledBody.Next
	map_conf.prv_pg = unmarshalledBody.Previous
	return nil
}

func getFromWeb(url string) ([]byte, error) {
	res, err := http.Get(url)
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 { //Check for error related status codes
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil { log.Fatal(err) }
	return body, err
}


func commandExit(map_conf *config, args []string) error {
        if len(args) >= 1 {
		fmt.Println("expected no arguments recieved ", len(args), "\n", args)
		return nil
	}
        fmt.Println("Closing the Pokedex... Goodbye!")
        os.Exit(0)
        return nil
}

func commandHelp(map_conf *config, args []string) error {
        if len(args) >= 1 {
		fmt.Println("expected no arguments recieved ", len(args), "\n", args)
		return nil
	}
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
