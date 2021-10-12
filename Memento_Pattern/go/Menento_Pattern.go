package Menento_Pattern

import "fmt"

type Text string

type TextEdit struct {
	text    Text
	Memento TextMemento
}

func (this *TextEdit) Append(s string) {
	this.Memento.Push(this.text)
	this.text = this.text + Text(s)
}
func (this *TextEdit) Revert() {
	this.text = this.Memento.Pop()
}

func (this *TextEdit) Show() {
	fmt.Println(this.text)
}

type TextMemento struct {
	history []Text
}

func (this *TextMemento) Push(t Text) {
	this.history = append(this.history, t)
}

func (this *TextMemento) Pop() Text {
	if len(this.history) == 0 {
		return ""
	}

	t := this.history[len(this.history)-1]
	this.history = this.history[:len(this.history)-1]
	return t
}
