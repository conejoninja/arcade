package main

import (
	"machine"
	"time"

	"github.com/conejoninja/arcade/tinyssd1306"
)

var (
	btnLeft  = machine.BUTTON_LEFT
	btnRight = machine.BUTTON_RIGHT
	display  *tinyssd1306.Device

	stopAnimate   bool
	lives         int8
	flipFlop      bool
	frogColumn    uint8
	frogRow       uint8
	grid          [6][16]uint8
	bounceLeft    bool
	bounceRight   bool
	bounceForward bool
	loopCounter   uint8
)

var bitmaps = [6][8]uint8{
	{0x83, 0xDC, 0x7A, 0x3F, 0x3F, 0x7A, 0xDC, 0x83},
	{0x3C, 0x7E, 0xD7, 0xB5, 0xAD, 0xBF, 0xFF, 0xED},
	{0xAD, 0xAD, 0xFF, 0xB7, 0xF5, 0xBF, 0xB7, 0xAD},
	{0xED, 0xBD, 0xC3, 0xBD, 0xA5, 0xBD, 0x42, 0x3C},
	{0x00, 0x1C, 0x22, 0x63, 0x7F, 0x7F, 0x22, 0x22},
	{0x22, 0x3E, 0x3E, 0x7F, 0x63, 0x63, 0x22, 0x1C},
}

func main() {
	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.SPEAKER.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.SPEAKER.Low()

	display = tinyssd1306.New(machine.P2, machine.P3)
	display.Configure()
	display.ClearScreen()

	for {
		playGame()
		display.FillScreen(0x00)
		time.Sleep(2 * time.Second)
	}
}

func playGame() {
	stopAnimate = false
	frogColumn = 8
	frogRow = 7
	flipFlop = true
	lives = 2
	loopCounter = 0

	initScreen()
	display.FillScreen(0x00)
	drawScreen()

	for lives >= 0 {
		loopCounter++
		if loopCounter >= 60 {
			loopCounter = 0
			moveBlocks()
			checkCollision()
			if !stopAnimate {
				drawScreen()
			}
		}

		handleInput()
		checkCollision()

		if stopAnimate {
			time.Sleep(200 * time.Millisecond)
			lives--
			stopAnimate = false
			frogColumn = 8
			frogRow = 7
			drawScreen()
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func handleInput() {
	l := btnLeft.Get()
	r := btnRight.Get()

	if l && r && !bounceForward {
		bounceForward = true
		if frogRow > 0 {
			frogRow--
		}
		drawScreen()
	} else if !l || !r {
		bounceForward = false
	}

	if l && !r && !bounceLeft {
		bounceLeft = true
		if frogColumn > 0 {
			frogColumn--
		}
		drawScreen()
	} else if !l {
		bounceLeft = false
	}

	if !l && r && !bounceRight {
		bounceRight = true
		if frogColumn < 15 {
			frogColumn++
		}
		drawScreen()
	} else if !r {
		bounceRight = false
	}
}

func checkCollision() {
	if frogRow > 0 && frogRow < 4 && grid[frogRow-1][frogColumn] == 0 {
		stopAnimate = true
	}
	if frogRow > 3 && frogRow < 7 && grid[frogRow-1][frogColumn] != 0 {
		stopAnimate = true
	}
	if frogRow == 0 {
		frogRow = 7
		frogColumn = 8
	}
}

func initScreen() {
	for r := 0; r < 6; r++ {
		for c := 0; c < 16; c++ {
			grid[r][c] = 0
		}
	}

	grid[0][0], grid[0][1], grid[0][2] = 2, 3, 4
	grid[0][6], grid[0][7], grid[0][8] = 2, 3, 4
	grid[0][12], grid[0][13], grid[0][14] = 2, 3, 4

	grid[1][2], grid[1][3], grid[1][4] = 2, 3, 4
	grid[1][7], grid[1][8], grid[1][9] = 2, 3, 4
	grid[1][12], grid[1][13], grid[1][14] = 2, 3, 4

	grid[2][1], grid[2][2], grid[2][3], grid[2][4] = 2, 3, 3, 4
	grid[2][8], grid[2][9], grid[2][10], grid[2][11] = 2, 3, 3, 4

	grid[3][0], grid[3][1] = 5, 6
	grid[3][8], grid[3][9] = 5, 6

	grid[4][3], grid[4][4] = 5, 6
	grid[4][9], grid[4][10] = 5, 6

	grid[5][1], grid[5][2] = 5, 6
	grid[5][11], grid[5][12] = 5, 6
}

func drawScreen() {
	for row := uint8(0); row < 6; row++ {
		inv := row < 3
		setPos(0, row+1)
		display.DataStart()
		for col := uint8(0); col < 16; col++ {
			if frogRow == row+1 && frogColumn == col {
				sendBlock(1, false)
			} else {
				sendBlock(grid[row][col], inv)
			}
		}
		display.DataStop()
	}

	if frogRow == 7 || frogRow == 0 {
		setPos(frogColumn*8, 7)
		display.DataStart()
		sendBlock(1, false)
		display.DataStop()
	}
}

func sendBlock(f uint8, inv bool) {
	if f == 0 {
		var b uint8
		if inv {
			b = 0xFF
		}
		display.DataByte(b)
		display.DataByte(b)
		display.DataByte(b)
		display.DataByte(b)
		display.DataByte(b)
		display.DataByte(b)
		display.DataByte(b)
		display.DataByte(b)
	} else {
		for i := 0; i < 8; i++ {
			b := bitmaps[f-1][i]
			if inv {
				b = ^b
			}
			display.DataByte(b)
		}
	}
}

func moveBlocks() {
	flipFlop = !flipFlop

	for row := 0; row < 6; row++ {
		goL := (row % 2) == 0
		if row >= 3 {
			goL = !goL
		}

		if frogRow > 0 && frogRow < 4 && frogRow == uint8(row)+1 {
			if goL {
				if frogColumn > 0 {
					frogColumn--
				} else {
					stopAnimate = true
				}
			} else if flipFlop && frogColumn < 15 {
				frogColumn++
			}
		}

		if goL {
			t := grid[row][0]
			for c := 0; c < 15; c++ {
				grid[row][c] = grid[row][c+1]
			}
			grid[row][15] = t
		} else if flipFlop {
			t := grid[row][15]
			for c := 15; c > 0; c-- {
				grid[row][c] = grid[row][c-1]
			}
			grid[row][0] = t
		}
	}
}

func setPos(x uint8, y uint8) {
	display.SendCommand(0xB0 + y)
	display.SendCommand(((x & 0xF0) >> 4) | 0x10)
	display.SendCommand((x & 0x0F) | 0x01)
}
