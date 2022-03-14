package main

import (
	"fmt"
	"log"
)

func fatal(err error, msg string, args ...interface{}) {
	m := fmt.Sprintf(msg, args...)
	log.Fatalf("%s: %s", m, err)
}
