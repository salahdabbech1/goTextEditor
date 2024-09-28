package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func print_message(column, row int, foregound, background termbox.Attribute, message string) {
	for _, character := range message {
		termbox.SetCell(column, row, character, foregound, background)
		column += runewidth.RuneWidth(character)
	}
}
func run_editor() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for {
		print_message(0, 0, termbox.ColorDefault, termbox.ColorDefault, "First time trying a text editor")
		termbox.Flush()
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
			termbox.Close()
			break
		}
	}
}
func main() {
	run_editor()
}
