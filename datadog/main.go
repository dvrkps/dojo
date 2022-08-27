package main

import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

func main() {
	tracer.Start()
	defer tracer.Stop()
}
