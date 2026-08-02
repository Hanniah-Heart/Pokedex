package main

type config struct {
        prv_pg string
        nxt_pg string
}

type cliCommand struct {
        name        string
        description string
        callback    func(*config) error
}
