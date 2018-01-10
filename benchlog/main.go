package main

import "log"

type nilLog struct {
	log *log.Logger
}

func (l *nilLog) Print(v ...interface{}) {
	if l.log == nil {
		return
	}
	l.Print(v...)
}
