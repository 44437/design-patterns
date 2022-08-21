package main

import "fmt"

func main() {
	var builderBaseComputer IBuilder
	var builderBaseComputerWithDisplay IBuilder
	var builderFullComputer IBuilder

	var directorOne *Director
	var directorTwo *Director

	builderBaseComputer = getBuilder("base")
	builderBaseComputerWithDisplay = getBuilder("base_with_display")
	builderFullComputer = getBuilder("full")

	directorOne = newDirector(builderBaseComputer)
	baseComputer := directorOne.buildComputer()
	fmt.Println(baseComputer)

	directorOne.setBuilder(builderBaseComputerWithDisplay)
	baseComputerWithDisplay := directorOne.buildComputer()
	fmt.Println(baseComputerWithDisplay)

	directorTwo = newDirector(builderFullComputer)
	fullComputer := directorTwo.buildComputer()
	fmt.Println(fullComputer)
}
