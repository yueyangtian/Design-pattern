package Prototype_Pattern

import (
	"fmt"
)

type inode interface {
	Clone() inode
	Print()
}

type File struct {
	name string
}

func (f *File) Clone() inode {
	return &File{name: f.name}
}
func (f *File) Print() {
	fmt.Println(f.name)
}

type Floder struct {
	name    string
	childen []inode
}

func (f *Floder) Clone() inode {
	res := &Floder{name: f.name}
	var clonechilden []inode
	for _, node := range f.childen {
		clonechilden = append(clonechilden, node)
	}
	res.childen = clonechilden
	return res
}
func (f *Floder) Print() {
	fmt.Printf("%s\n", f.name)
	for _, node := range f.childen {
		fmt.Print("		")
		node.Print()
	}
}
