package main

import "fmt"

func GetGun(gunType string) (Gun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
