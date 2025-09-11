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
	oldPlayer uint8 = 61

	shootsX [3]uint8 = [3]uint8{0, 0, 0}
	shootsY [3]uint8 = [3]uint8{0, 0, 0}
)

func main() {

	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	display = tinyssd1306.New(machine.P2, machine.P3)
	display.Configure()
	display.ClearScreen()

	shootsY[0] = 55
	shootsX[0] = player

	display.SetPos(61, 0)
	display.DataStart()
	display.DataByte(0xFF)
	display.DataByte(0xF0)
	display.DataByte(0x78)
	display.DataByte(0x3C)
	display.DataByte(0x1E)
	display.DataByte(0x0F)
	display.DataStop()

	for {
		if btnLeft.Get() && player > 0 {
			player -= 3
		}
		if btnRight.Get() && player < 120 {
			player += 3
		}

		if shootsY[0] == 0 {
			display.SetPos(shootsX[0], 0)
			display.DataStart()
			display.DataByte(0x00)
			display.DataStop()

			shootsY[0] = 55
			shootsX[0] = player
		}
		shootsY[0]--
		DrawShoot(0)

		if oldPlayer != player {

			ClearPlayer(oldPlayer)
			DrawPlayer(player)
			oldPlayer = player
		}

		//time.Sleep(50 * time.Millisecond)

	}
}

func DrawPlayer(x uint8) {
	display.SetPos(x, 63)
	display.DataStart()
	display.DataByte(0xC0)
	display.DataByte(0x60)
	display.DataByte(0x30)
	display.DataByte(0x7C)
	display.DataByte(0x30)
	display.DataByte(0x60)
	display.DataByte(0xC0)
	display.DataStop()
}

func ClearPlayer(x uint8) {
	display.SetPos(x, 63)
	display.DataStart()
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataStop()
}

func DrawShoot(index uint8) {
	i := shootsY[index] % 8
	if i >= 5 && shootsY[index] < 56 {
		display.SetPos(shootsX[index], 8*(shootsY[index]/8)+8)
		display.DataStart()
		display.DataByte(0x0F >> (8 - i))
		display.DataStop()
	}
	if i == 4 && shootsY[index] < 56 {
		display.SetPos(shootsX[index], 8*(shootsY[index]/8)+8)
		display.DataStart()
		display.DataByte(0x00)
		display.DataStop()
	}
	display.SetPos(shootsX[index], 8*(shootsY[index]/8))
	display.DataStart()
	display.DataByte(0x0F << i)
	display.DataStop()
}
