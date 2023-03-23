package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

var err = errors.New("it's friday, I'm done")

func TestFirstCase(t *testing.T) {
	ctx := context.Background()

	ctxWithTimeout, _ := context.WithTimeout(ctx, time.Microsecond)
	time.Sleep(time.Millisecond)
	fmt.Println("err: ", ctxWithTimeout.Err().Error())

	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	fmt.Println("err: ", ctxWithCancel.Err().Error())
}

func TestSecondCase(t *testing.T) {
	ctx := context.Background()

	ctxWithCancel, cancel := context.WithCancelCause(ctx)
	cancel(err)
	fmt.Println("cause: ", context.Cause(ctxWithCancel))
	fmt.Println("err: ", ctxWithCancel.Err().Error())
}
