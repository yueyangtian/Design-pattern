package Adapter_Pattern

import (
	"testing"
)

func TestAdapterPattern(t *testing.T) {
	client := &Client{}
	mac := &Mac{}
	client.insertInSquarePort(mac)

	win := &Windows{}
	windowsadapter := &WindowsAdapter{windows: win}
	client.insertInSquarePort(windowsadapter)
}
