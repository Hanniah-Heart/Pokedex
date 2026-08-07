package main
import (
	"time"
	pokecache "github.com/Hanniah-Heart/Pokedexcli/internal/pokecache"
)

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

type config struct {
        prv_pg string
        nxt_pg string
	reapingInterval time.Duration
	CacheAddress *pokecache.Cache
	Pokedex map[string]pokemon
}

type cliCommand struct {
        name        string
        description string
        callback    func(*config, []string) error
}
