package main

import "fmt"

type Folder struct {
	Components []Component
	Name       string
	Size       string
	Format     string
}

func (f *Folder) Delete() {
	*f = Folder{}
}

func (f *Folder) Search(what string) {
	fmt.Printf("Searching for keyword %s in FOLDER %s\n", what, f.Name)
	for _, component := range f.Components {
		component.Search(what)
	}
}

func (f *Folder) Rename(new string) {
	f.Name = new
}

func (f *Folder) AddComponent(component Component) {
	f.Components = append(f.Components, component)
}
