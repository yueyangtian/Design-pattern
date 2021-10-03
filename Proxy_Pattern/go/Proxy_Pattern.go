package Proxy_Pattern

import "fmt"

type image interface {
	display()
}

type RealImage struct {
	Name string
}

func (this *RealImage) display() {
	this.load()
	fmt.Println("image displayed")
}

func (this *RealImage) load() {
	fmt.Printf("image %s loading \n", this.Name)
}

type ProxyImage struct {
	image RealImage
}

func (this *ProxyImage) display() {
	this.image.display()
}

func (this *ProxyImage) proxy(name string) {
	this.image.Name = name
}
