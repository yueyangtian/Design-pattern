package DesginPattern

import (
	"sync"
)

type Singleton struct{}

var instance *Singleton
var hungryinstance *Singleton
var mu sync.Mutex

//double check lock for lazy
func GetInstance() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}
func GetInstanceHungry() *Singleton {
	return hungryinstance
}

//hungry model
func init() {
	hungryinstance = &Singleton{}
}
