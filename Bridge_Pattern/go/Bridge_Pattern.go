package Bridge_Pattern

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
type hp_windows struct {
}

func (p *hp_windows) printFile() {
	fmt.Println("OS:windows, Printing by a HP Printer")
}

type hp_linux struct {
}

func (p *hp_linux) printFile() {
	fmt.Println("OS:linux, Printing by a HP Printer")
}

//ms
type suf_windows struct {
}

func (s *suf_windows) printFile() {
	fmt.Println("OS:windows, Print by a Suf Printer")
}
