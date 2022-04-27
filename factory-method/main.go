package main

import "fmt"

func main() {
	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g Gun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
