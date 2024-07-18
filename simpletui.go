package simpletui

import "fmt"

func MoveCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH]", x, y)
}

func Label(x int, y int, text string) {
	MoveCursor(x, y)
	fmt.Printf("`%s`\n", text)
}
