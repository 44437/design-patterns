package main

import "fmt"

type VideoPlayer struct{}

func (v *VideoPlayer) Play(work string) {
	fmt.Printf("%s is playing\n", work)
}
