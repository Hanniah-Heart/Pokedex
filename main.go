package main
import (
	"fmt"
	"bufio"
	"os"
	"time"
	pokecache "github.com/Hanniah-Heart/Pokedexcli/internal/pokecache"
)

func main() {
	interval, _ := time.ParseDuration("5s")
	map_conf := config {
		prv_pg: "",
		nxt_pg: "https://pokeapi.co/api/v2/location-area/",
		reapingInterval: interval,
	}
	newCache := pokecache.NewCache(map_conf.reapingInterval)
	map_conf.CacheAddress = newCache
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
		arguments := (cleanInput(scanner.Text()))[1:]
		cmd, exists := commandList[userCommand]
		if exists == false {
			fmt.Print("Unknown command")
		} else {
			err := cmd.callback(&map_conf, arguments)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}
