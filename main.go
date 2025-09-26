package main

import (
	"machine"
	"time"

	"github.com/conejoninja/arcade/tinyssd1306"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var (
	sda_pin machine.Pin = machine.P2
	scl_pin machine.Pin = machine.P3

	btnLeft  machine.Pin = machine.BUTTON_LEFT
	btnRight machine.Pin = machine.BUTTON_RIGHT
	speaker  machine.Pin = machine.SPEAKER

	display *tinyssd1306.Device

	snake                   [128]uint8
	snakeLength             uint8
	snakeDirection          uint8
	snakeHead               uint8
	apple                   uint8
	i                       uint8
	bounceLeft, bounceRight bool
)

func main() {

	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})

	display = tinyssd1306.New(machine.P2, machine.P3)
	display.Configure()
	display.ClearScreen()

	startGame()

	for {
		if btnLeft.Get() {
			if !bounceLeft {
				bounceLeft = true
				if snakeDirection == 0 {
					snakeDirection = 3
				} else {
					snakeDirection--
				}
			}
		} else {
			bounceLeft = false
		}
		if btnRight.Get() {
			if !bounceRight {
				bounceRight = true
				if snakeDirection == 3 {
					snakeDirection = 0
				} else {
					snakeDirection++
				}
			}
		} else {
			bounceRight = false
		}

		moveSnake()
	}
}

func moveSnake() {
	snakeHead = snake[0]
	switch snakeDirection {
	case RIGHT:
		if snakeHead%16 == 15 {
			snake[0] = snakeHead - 15
		} else {
			snake[0] = snakeHead + 1
		}
		break
	case UP:
		if snakeHead < 16 {
			snake[0] = snakeHead + 128 - 16
		} else {
			snake[0] = snakeHead - 16
		}
		break
	case LEFT:
		if snakeHead%16 == 0 {
			snake[0] = snakeHead + 15
		} else {
			snake[0] = snakeHead - 1
		}
		break
	case DOWN:
		if snakeHead > 112 {
			snake[0] = snakeHead % 16
		} else {
			snake[0] = snakeHead + 16
		}
		break
	}

	drawRect(snake[0])
	if apple == snake[0] {
		snakeLength++
		apple += 7 * snakeLength
		if apple > 127 {
			apple -= 128
		}

		drawApple(apple)
	} else {
		for i = snakeLength - 1; i > 2; i-- {
			if snake[i] == snake[0] {
				startGame()
				return
			}
		}
		clearRect(snake[snakeLength-1])
	}
	for i = snakeLength - 1; i > 0; i-- {
		snake[i] = snake[i-1]
	}
}

func startGame() {
	song()
	display.FillScreen(0xFF)
	display.FillScreen(0x00)
	snakeLength = 4
	snakeDirection = RIGHT

	snake[0] = 34

	apple = 100
	drawApple(apple)
}

func drawRect(i uint8) {
	display.SetPos(8*(i%16), 8*(i/16))
	display.DataStart()
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataByte(0xFF)
	display.DataStop()
}

func clearRect(i uint8) {
	display.SetPos(8*(i%16), 8*(i/16))
	display.DataStart()
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataByte(0x00)
	display.DataStop()
}

func drawApple(i uint8) {
	display.SetPos(8*(i%16), 8*(i/16))
	display.DataStart()
	display.DataByte(0xFF)
	display.DataByte(0x81)
	display.DataByte(0x81)
	display.DataByte(0x81)
	display.DataByte(0x81)
	display.DataByte(0x81)
	display.DataByte(0x81)
	display.DataByte(0xFF)
	display.DataStop()
}

func beep(count, delay int) {
	for i := 0; i <= count; i++ {
		speaker.High()
		time.Sleep(time.Duration(delay) * time.Microsecond)
		speaker.Low()
		time.Sleep(time.Duration(delay) * time.Microsecond)
	}
}

func song() {
	beep(50, 250)
	time.Sleep(100 * time.Millisecond)
	beep(150, 150)
	time.Sleep(100 * time.Millisecond)
	beep(250, 50)
}
