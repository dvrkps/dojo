// https://play.golang.org/p/X6jvh3zjY7

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func NaiveRetry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return NaiveRetry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

// Retry will run fn and return immediately if it succeeds, otherwise it will
// retry while respecting backoff until the given context is canceled.
//
// if err := Retry(ctx, func() error { return nil }); err != nil {
// 	fmt.Println(`unexpected err:`, err)
// }
func Retry(ctx context.Context, fn func() error) (err error) {
	last, err := time.Now(), fn()
	for attempts := 0; err != nil; attempts++ {
		sleep := BackoffSince(attempts, time.Since(last))
		select {
		case <-time.After(sleep):
		case <-ctx.Done():
			return err
		}
		last, err = time.Now(), fn()
	}
	return
}

func main() {
	{
		fmt.Println(`Exponential Backoff:`)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		for i := 0; i < 20; i++ {
			go func() {
				begin, attempts, elapsed := time.Now(), 0, time.Duration(0)

				fmt.Printf("Retry(%v)\n", ctx)
				err := Retry(ctx, func() error {
					attempts, elapsed = attempts+1, time.Since(begin)
					err := fmt.Errorf("  fail attempt #%d (+%v)", attempts, elapsed)
					if elapsed > time.Second*8 {
						fmt.Println("  success!")
						return nil
					}
					fmt.Println(err)
					return err
				})
				if err != nil {
					fmt.Println(`unexpected err:`, err)
				}
			}()
		}
		<-ctx.Done()
	}
	{
		fmt.Println(`Naive Backoff:`)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		for i := 0; i < 20; i++ {
			go func() {
				begin, attempts, elapsed := time.Now(), 0, time.Duration(0)

				fmt.Printf("Retry(%v)\n", ctx)
				err := NaiveRetry(15, time.Second, func() error {
					attempts, elapsed = attempts+1, time.Since(begin)
					err := fmt.Errorf("  fail attempt #%d (+%v)", attempts, elapsed)
					if elapsed > time.Second*8 {
						fmt.Println("  success!")
						return nil
					}
					fmt.Println(err)
					return err
				})
				if err != nil {
					fmt.Println(`unexpected err:`, err)
				}
			}()
		}
		<-ctx.Done()
	}
}

var backoffTab = make([]time.Duration, 15)

const (
	multiplier = 1.40
)

func init() {
	cur := float64(250 * time.Millisecond)
	for i, max := 0, len(backoffTab); i < max; i++ {
		cur *= multiplier
		backoffTab[i] = time.Duration(cur)
	}
}

// Backoff will return the amount of time before retrying a operation based on
// the current resource limitations.
func Backoff(attempts int) time.Duration {
	return BackoffSince(attempts, 0)
}

// BackoffSince will return the amount of time before retrying a operation based
// on the current resource limitations taking the elapsed duration since the
// last retry into consideration if it is non-zero.
func BackoffSince(attempts int, last time.Duration) time.Duration {
	if last != 0 {
		max := backoffTab[len(backoffTab)-1]
		if last > max {
			attempts = 0
		}
	}
	if attempts >= len(backoffTab) {
		attempts = len(backoffTab) - 1
	}
	return Jitter(backoffTab[attempts])
}

// Jitter will add jitter to a time.Duration.
func Jitter(d time.Duration) time.Duration {
	const jitter = 0.30
	jit := 1 + jitter*(rand.Float64()*2-1)
	return time.Duration(jit * float64(d))
}
