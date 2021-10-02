package Facade_Pattern

import "fmt"

type HardDrive struct{}

func (*HardDrive) load() {
	fmt.Println("hard drive loading")
}

type Memory struct{}

func (*Memory) read() {
	fmt.Println("memory reading")
}

type CPU struct{}

func (*CPU) execute() {
	fmt.Println("cpu executing")
}

type Computer struct {
	hardDrive HardDrive
	memory    Memory
	cpu       CPU
}

func (this *Computer) start() {
	this.hardDrive.load()
	this.cpu.execute()
	this.memory.read()
}
