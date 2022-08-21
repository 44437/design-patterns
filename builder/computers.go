package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	Computer interface{}

	BaseComputer struct {
		Ram       string
		CPU       string
		HardDrive string
	}

	BaseComputerWithDisplay struct {
		Ram         string
		CPU         string
		HardDrive   string
		Display     string
		GraphicCard string
	}

	FullComputer struct {
		Ram         string
		CPU         string
		HardDrive   string
		Display     string
		GraphicCard string
		Keyboard    string
		Mouse       string
	}
)

func (b BaseComputer) String() string {
	marshall, err := json.Marshal(b)
	if err != nil {
		log.Panicln(err)
	}
	return fmt.Sprintf(string(marshall))
}

func (b BaseComputerWithDisplay) String() string {
	marshall, err := json.Marshal(b)
	if err != nil {
		log.Panicln(err)
	}
	return fmt.Sprintf(string(marshall))
}

func (f FullComputer) String() string {
	marshall, err := json.Marshal(f)
	if err != nil {
		log.Panicln(err)
	}
	return fmt.Sprintf(string(marshall))
}
