package main
import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
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

func commandMap() error {
	response, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	if response.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	type result struct {
		Name	string
		Url	string
	}
	type mappage struct {
		Count		int
		Next		string
		Previous	string
		Results		[]result
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
}// This seems to work, but now how do you handle moving between pages?

func commandMapOut(*config) error {
	
	response, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	if response.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	type result struct {
		Name	string
		Url	string
	}
	type mappage struct {
		Count		int
		Next		string
		Previous	string
		Results		[]result
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


func commandMapb() error {
	/*if config.prv_page == "" {
		return fmt.Errorf("you're on the first page")
	}*/
	response, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	if response.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	type result struct {
		Name	string
		Url	string
	}
	type mappage struct {
		Count		int
		Next		string
		Previous	string
		Results		[]result
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

type config struct {
	prv_page string
	nxt_page string
}

type cliCommand struct {
	name        	string
	description 	string
	callback    	func(*config) error
}

func getCommandList() map[string]cliCommand {
	commandList := map[string]cliCommand{
		"help": {
			name: "help",
			description: "Describe how to use the Pokedex",
			callback: commandHelp(*mapconfig),
		}, "exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit(),
		}, "map": {
			name: "map",
			description: "Opens next map page",
			callback: commandMap(),
		}, "mapb": {
			name: "mapb",
			description: "Opens previous map page",
			callback: commandMap(),
		}, "mapout": {
			name: "mapout",
			description: "Should not exist. Part of debugging.",
			callback: commandMap(config_pointer),
		},
	}
	return commandList
}

func main() {
	// INITIALIZE A CONFIG!!!
	mapconfig := config {
		prv_page: "",
		nxt_page: "",
	}
	config_pointer := &mapconfig
	fmt.Printf(config_pointer.prv_page)
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
			err := cmd.callback(config_pointer)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}
