package main

type Component interface {
	Delete()
	Search(what string)
	Rename(new string)
}
