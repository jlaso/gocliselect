package main

import "go-cli-select"

func main() {
	menu := go_cli_select.NewMenu("", "Select an option")

	menu.AddItem("Option 1", "option1")
	menu.AddItem("Option 2", "option2")
	menu.AddItem("Option 3", "option3")

	menu.Display()
}