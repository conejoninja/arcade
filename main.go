package main

import (
	"machine"
	"time"

	"github.com/conejoninja/arcade/tinyssd1306"
)

var (
	sda_pin machine.Pin = machine.P2
	scl_pin machine.Pin = machine.P3
)

var display *tinyssd1306.Device

func main() {

	display = tinyssd1306.New(machine.P2, machine.P3)
	display.Configure()
	display.ClearScreen()
	for {
		display.FillScreen(0x11)
		time.Sleep(500 * time.Millisecond)
		display.FillScreen(0x00)
		time.Sleep(500 * time.Millisecond)

		display.SetPixel(0, 0, true)
		display.SetPixel(8, 8, true)
		display.SetPixel(16, 16, true)
		display.SetPixel(24, 24, true)

		time.Sleep(500 * time.Millisecond)

	}
}
