package main

import (
	"machine"

	"github.com/conejoninja/arcade/tinyssd1306"
)

var (
	sda_pin machine.Pin = machine.P2
	scl_pin machine.Pin = machine.P3

	btnLeft  machine.Pin = machine.BUTTON_LEFT
	btnRight machine.Pin = machine.BUTTON_RIGHT

	display *tinyssd1306.Device

	player    uint8 = 60
	oldPlayer uint8 = 60
)

func main() {

	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	display = tinyssd1306.New(machine.P2, machine.P3)
	display.Configure()
	display.ClearScreen()

	for {
		if btnLeft.Get() && player > 0 {
			player--
		}
		if btnRight.Get() && player < 116 {
			player++
		}

		if oldPlayer != player {
			display.SetPixel(oldPlayer, 48, false)
			display.SetPixel(player, 48, true)
			oldPlayer = player
		}

		//time.Sleep(50 * time.Millisecond)

	}
}
