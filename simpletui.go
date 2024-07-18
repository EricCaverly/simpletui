package simpletui

import (
	"fmt"
	"strings"
)

func MoveCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func Label(x int, y int, text string) {
	MoveCursor(x, y)
	fmt.Printf("%s", text)
}

func TextBox(x int, y int, text string, borderStyle int) {

	// Split the text by newlines
	split_text := strings.Split(text, "\n")
	max_len := 0

	// Find the longest line
	for _, v := range split_text {
		if len(v) > max_len {
			max_len = len(v)
		}
	}

	// Create the line to be used for the top and bottom borders
	line := strings.Repeat("─", max_len)

	// Draw the top border
	MoveCursor(x-1, y-1)
	fmt.Printf("┌%s┐\n", line)

	// Draw the left border, text, and right border
	for i, v := range split_text {
		MoveCursor(x-1, y+i)
		fmt.Printf("│%-*s│\n", max_len, v)
	}

	// Draw the bottom border
	MoveCursor(x-1, y+len(split_text))
	fmt.Printf("└%s┘", line)
}
