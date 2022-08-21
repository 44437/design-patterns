package main

import (
	"fmt"
)

type IBuilder interface {
	buildRam()
	buildCPU()
	buildHardDrive()
	buildDisplay()
	buildGraphicCard()
	buildKeyboard()
	buildMouse()
	getComputer() Computer
}

func getBuilder(computerType string) IBuilder {
	switch computerType {
	case "base":
		return newBaseComputerBuilder()
	case "base_with_display":
		return newBaseComputerWithDisplayBuilder()
	case "full":
		return newFullComputerBuilder()
	}
	panic(fmt.Sprintf("%s not found", computerType))
}
