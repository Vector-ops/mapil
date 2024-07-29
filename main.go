package main

import (
	"github.com/vector-ops/mapil/cmd"
	"github.com/vector-ops/mapil/store"
)

func main() {

	store := store.NewStore()
	store.Init()

	cmd.Execute(store)
}
