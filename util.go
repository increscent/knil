package main

import (
	"math/rand"
	"time"
)

func uuid() uint64 {
	return rand.New(
		rand.NewSource((int64)(time.Now().Nanosecond()))).Uint64()
}