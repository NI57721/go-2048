package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	// defer term.Close()

	g := Game{}

	g.draw()
	g.draw()

main_loop:
	for {
		term.Sync()
		fmt.Println(g.String())
		fmt.Println("Press ↑↓→← or kjlh. Press ESC/q to quit.")
		switch ev := term.PollEvent(); ev.Type {
		case term.EventError:
			panic(ev.Err)
		case term.EventKey:
			switch {
			case ev.Key == term.KeyArrowUp || ev.Ch == 'k':
				g.drawIfAbleToLean("up")
			case ev.Key == term.KeyArrowDown || ev.Ch == 'j':
				g.drawIfAbleToLean("down")
			case ev.Key == term.KeyArrowRight || ev.Ch == 'l':
				g.drawIfAbleToLean("right")
			case ev.Key == term.KeyArrowLeft || ev.Ch == 'h':
				g.drawIfAbleToLean("left")
			case ev.Ch == 'd':
				g.draw()
			case ev.Key == term.KeyEsc || ev.Ch == 'q':
				break main_loop
			default:
				// 	term.Sync()
				// 	fmt.Println("ASCII: ", ev.Ch)
				// 	fmt.Println("key: ", ev.Key)
			}
		}
	}
}
