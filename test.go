package main

import (
	"fmt"
	"time"
)

func main() {
	var filename string = "test.go"
	var t1 = time.Now()
	fmt.Printf("It is month: %d, day: %d, hour: %d,  minute: %d in filename: %s\n", t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), filename)

	var t2 = time.Since(t1)

	fmt.Printf("Its been %2.f minutes %2.f seconds and %2.d milliseconds since the first statement\n", t2.Minutes(), t2.Seconds(), t2.Milliseconds())
}
