package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const (
	serviceName    = "myservice"
	serviceVersion = "v0.1.2"
)

func main() {
	tracer.Start(
		tracer.WithService(serviceName),
		tracer.WithServiceVersion(serviceVersion),
		tracer.WithTraceEnabled(true),
	)
	defer tracer.Stop()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := run(ctx)
	if err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("run: %v", err)
		}
	}
}

func run(ctx context.Context) error {
	span := tracer.StartSpan(
		"run",
		tracer.ServiceName(serviceName),
		tracer.ResourceName("resource_name"),
	)

	ctx = tracer.ContextWithSpan(ctx, span)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
			first(ctx)

		}
	}
}

func first(ctx context.Context) {
	println("first")
	span, ctx := tracer.StartSpanFromContext(ctx, "first")
	defer span.Finish(tracer.FinishTime(time.Now()))

	err := second(ctx)
	if err != nil {
		span.Finish(tracer.WithError(err))
	}
}

func second(ctx context.Context) error {
	span, _ := tracer.StartSpanFromContext(ctx, "second")
	defer span.Finish(tracer.FinishTime(time.Now()))
	x := rand.Intn(2)
	println("second",x)
	if x == 1 {
		span.Finish(tracer.WithError(errors.New("ups")))
	}
	return nil
}
