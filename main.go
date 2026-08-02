package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	map_conf := config {
		prv_pg: "",
		nxt_pg: "https://pokeapi.co/api/v2/location-area/",
	}
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
			err := cmd.callback(&map_conf)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}
