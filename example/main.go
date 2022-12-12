package main

import (
	"fmt"

	"github.com/rrossmiller/gocliselect"
)

func main() {
	menu := gocliselect.NewMenu("Chose a colour")
	menu.VimKeys = true

	menu.AddItem("Red", "red")
	menu.AddItem("Blue", "blue")
	menu.AddItem("Green", "green")
	menu.AddItem("Yellow", "yellow")
	menu.AddItem("Cyan", "cyan")

	choice := menu.Display()

	fmt.Printf("Choice: %s\n", choice)
}
