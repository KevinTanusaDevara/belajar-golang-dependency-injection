package simple

import "fmt"

type File struct {
	Name string
}

func (f File) Closer() {
	fmt.Println("Close File", f.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Closer()
	}
}