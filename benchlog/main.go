package main

import "log"

type nilLog struct {
	log *log.Logger
}

func (l *nilLog) Print(v ...interface{}) {
	if l.log == nil {
		return
	}
	l.log.Print(v...)
}

type normalLog struct {
	log *log.Logger
}

func (l *normalLog) Print(v ...interface{}) {
	l.log.Print(v...)
}

func main() {}
