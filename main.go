package main

import (
	"github.com/vector-ops/mapil-cli/cmd"
	"github.com/vector-ops/mapil-cli/store"
)

func main() {

	store := store.NewStore()
	store.Init()

	cmd.Execute(store)
}
