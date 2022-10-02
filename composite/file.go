package main

import "fmt"

type File struct {
	Name   string
	Size   string
	Format string
}

func (f *File) Delete() {
	*f = File{}
}

func (f *File) Search(what string) {
	fmt.Printf("Searching for keyword %s in FILE %s\n", what, f.Name)
}

func (f *File) Rename(new string) {
	f.Name = new
}

func (f *File) String() string {
	return fmt.Sprintf("File Name :%s \t File Size:%s \t File Format:%s", f.Name, f.Size, f.Format)
}
