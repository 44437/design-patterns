package main

import "fmt"

func main() {
	var file1 *File
	var file2 *File
	var folder *Folder

	file1 = &File{
		Name:   "design.jpeg",
		Size:   "300Kb",
		Format: "JPEG",
	}

	file2 = &File{
		Name:   "hello.png",
		Size:   "140Kb",
		Format: "PNG",
	}

	folder = &Folder{
		Components: []Component{
			file1,
			file2,
		},
		Name:   "my_pics",
		Size:   "440",
		Format: "Folder",
	}

	fmt.Println(folder)
	file1.Search("test")
	folder.Search("test")

	file1.Delete()
	fmt.Println(file1)
}
