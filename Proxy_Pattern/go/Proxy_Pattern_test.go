package Proxy_Pattern

import "testing"

func TestProxyPattern(t *testing.T) {
	p := ProxyImage{}
	p.proxy("1.jpg")
	p.display()
}
