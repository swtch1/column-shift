package main

import (
	"fmt"
	"log"
)

func info(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func fatal(err error, msg string, args ...interface{}) {
	m := fmt.Sprintf(msg, args...)
	log.Fatalf("%s: %s", m, err)
}
