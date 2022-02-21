package main

import (
	"fmt"
	"strings"
)

func main() {
	realPresident := NewRealPresident()
	presidency := NewPresidency(realPresident)
	citizen := NewCitizen(presidency)

	fmt.Println(strings.Repeat("*", 50))
	citizen.TellTrouble("help please")
	fmt.Println(strings.Repeat("*", 50))
	citizen.TellTrouble("help")
	fmt.Println(strings.Repeat("*", 50))
	citizen.WantJob("I want a job")
	fmt.Println(strings.Repeat("*", 50))

}
