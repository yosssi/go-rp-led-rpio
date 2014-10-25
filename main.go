package main

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(17)

	pin.Output()

	pin.High()

	time.Sleep(3 * time.Second)

	pin.Low()
}
