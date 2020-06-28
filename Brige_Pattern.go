package DesginPattern

import "fmt"

type printer interface {
	printFile()
}
type computer interface {
	print()
	setPrinter(printer)
}

//linux
type linux struct {
	printer printer
}

func (l *linux) print() {
	fmt.Println("Print request for linux")
	l.printer.printFile()
}

func (l *linux) setPrinter(p printer) {
	l.printer = p
}

//windows

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}
func (w *windows) setPrinter(p printer) {
	w.printer = p
}

//hp
type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}

//ms
type suf struct {
}

func (s *suf) printFile() {
	fmt.Println("Print by a Suf Printer")
}
