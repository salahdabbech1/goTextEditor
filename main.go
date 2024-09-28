package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

var ROWS, COLS int
var offsetX, offsetY int

var text_buffer = [][]rune{
	{'h', 'e', 'l', 'l', 'o'},
	{'w', 'o', 'r', 'l', 'd'},
}

func display_text_buffer() {
	var row, col int
	for row = 0; row < ROWS; row++ {
		text_buffer_row := row + offsetY
		for col = 0; col < COLS; col++ {
			text_buffer_col := col + offsetX
			if text_buffer_row >= 0 && text_buffer_row < len(text_buffer) && text_buffer_col < len(text_buffer[text_buffer_row]) {
				if (text_buffer[text_buffer_row][text_buffer_col]) != '\t' {
					termbox.SetChar(col, row, text_buffer[text_buffer_row][text_buffer_col])
				} else {
					termbox.SetCell(col, row, rune(' '), termbox.ColorDefault, termbox.ColorGreen)
				}
			} else if row+offsetY > len(text_buffer)-1 {
				termbox.SetCell(0, row, rune('*'), termbox.ColorCyan, termbox.ColorDefault)
				//termbox.SetChar(col, row, rune('\n'))
			}

		}
	}

}

/*
	func print_message(col, row int, foregound, background termbox.Attribute, message string) {
		for _, character := range message {
			termbox.SetCell(col, row, character, foregound, background)
			col += runewidth.RuneWidth(character)
		}
	}
*/
func run_editor() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for {
		ROWS, COLS = termbox.Size()
		ROWS--
		if COLS < 80 {
			COLS = 80
		}
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		display_text_buffer()
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
