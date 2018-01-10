package main

import "log"

type nilLog struct {
	log *log.Logger
}

func (l *nilLog) Print(v ...interface{}) {
	if l == nil {
		return
	}
	l.Print(v...)
}
