package main

import "fmt"

type Display struct{}

func (m *Display) Show(work string) {
	fmt.Printf("%s is showing\n", work)
}
