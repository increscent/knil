package main

import (
	"log"
	"math/rand"
	"os/exec"
	"regexp"
	"time"
)

func uuid() uint64 {
	return rand.New(
		rand.NewSource(int64(time.Now().Nanosecond()))).Uint64()
}

func sysuuid() string {
	out, err := exec.Command("uuidgen").Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(regexp.MustCompilePOSIX("[0-9,a-z,-]*").Find(out))
}
