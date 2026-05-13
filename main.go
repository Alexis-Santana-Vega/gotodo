package main

import (
	"github.com/Alexis-Santana-Vega/gotodo/todo"
	"github.com/Alexis-Santana-Vega/gotodo/ui"
)

func main() {
	store := todo.New()
	ui.Run(store)
}
