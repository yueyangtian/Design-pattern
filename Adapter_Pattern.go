package DesginPattern

import (
	"fmt"
)

type Computer interface {
	insertInSquarePort()
}

type Mac struct{}

func (m *Mac) insertInSquarePort() {
	fmt.Println("Mac insert square port")
}

type Windows struct{}

func (w *Windows) insertInCriclePort() {
	fmt.Println("Windows insert cricle port")
}

type WindowsAdapter struct {
	windows *Windows
}

func (w *WindowsAdapter) insertInSquarePort() {
	w.windows.insertInCriclePort()
}

type Client struct{}

func (c *Client) insertInSquarePort(com Computer) {
	com.insertInSquarePort()
}
