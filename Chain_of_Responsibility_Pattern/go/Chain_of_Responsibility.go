package Chain_of_Responsibility

import "fmt"

type Logger interface {
	write(message string)
	SetNextLog(Logger)
	Next() Logger
	Level() int
}

type Log struct {
	LoggerChain Logger
}

func (this *Log) Log(level int, message string) {
	c := this.LoggerChain
	for c != nil {
		if level != c.Level() {
			continue
		}
		c.write(message)
		c = c.Next()
	}
}

type Debug struct {
	level      int
	nextLogger Logger
}

func (this *Debug) write(message string) {
	fmt.Printf("[DEBUG] %s\n", message)
}
func (this *Debug) SetNextLog(l Logger) {
	this.nextLogger = l
}
func (this *Debug) Level() int {
	return this.level
}
func (this *Debug) Next() Logger {
	return this.nextLogger
}

type Info struct {
	level      int
	nextLogger Logger
}

func (this *Info) write(message string) {
	fmt.Printf("[INFO] %s\n", message)
}
func (this *Info) SetNextLog(l Logger) {
	this.nextLogger = l
}
func (this *Info) Level() int {
	return this.level
}
func (this *Info) Next() Logger {
	return this.nextLogger
}

type Error struct {
	level      int
	nextLogger Logger
}

func (this *Error) write(message string) {
	fmt.Printf("[ERROR] %s\n", message)
}
func (this *Error) SetNextLog(l Logger) {
	this.nextLogger = l
}
func (this *Error) Level() int {
	return this.level
}
func (this *Error) Next() Logger {
	return this.nextLogger
}
func GetLogger() *Log {
	eLog := Error{level: 1}
	iLog := Info{level: 2}
	dLog := Debug{level: 3}

	dLog.SetNextLog(&iLog)
	iLog.SetNextLog(&eLog)

	l := Log{LoggerChain: &eLog}

	return &l
}
