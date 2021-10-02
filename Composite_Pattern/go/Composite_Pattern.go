package Composite_Pattern

import (
	"fmt"
)

type Inode interface {
	search(string)
}

type ComFloder struct {
	nodes []Inode
	name  string
}

func (f *ComFloder) search(name string) {
	fmt.Printf("search file from folder [%s]\n", f.name)
	for _, n := range f.nodes {
		n.search(name)
	}
}

func (f *ComFloder) Add(n Inode) {
	f.nodes = append(f.nodes, n)
}

type ComLeaf struct {
	name string
}

func (f *ComLeaf) search(name string) {
	if f.name == name {
		fmt.Printf("find success\n")
	}
}
