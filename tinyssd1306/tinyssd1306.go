package tinyssd1306

import (
	"machine"
	"time"
)

const (
	ADDRESS = 0x3C

	SETCONTRAST         = 0x81
	DISPLAYALLON_RESUME = 0xA4
	DISPLAYALLON        = 0xA5
	NORMALDISPLAY       = 0xA6
	INVERTDISPLAY       = 0xA7
	DISPLAYOFF          = 0xAE
	DISPLAYON           = 0xAF
	SETDISPLAYOFFSET    = 0xD3
	SETCOMPINS          = 0xDA
	SETVCOMDETECT       = 0xDB
	SETDISPLAYCLOCKDIV  = 0xD5
	SETPRECHARGE        = 0xD9
	SETMULTIPLEX        = 0xA8
	SETLOWCOLUMN        = 0x00
	SETHIGHCOLUMN       = 0x10
	SETSTARTLINE        = 0x40
	MEMORYMODE          = 0x20
	COLUMNADDR          = 0x21
	PAGEADDR            = 0x22
	COMSCANINC          = 0xC0
	COMSCANDEC          = 0xC8
	SEGREMAP            = 0xA0
	CHARGEPUMP          = 0x8D
	EXTERNALVCC         = 0x1
	SWITCHCAPVCC        = 0x2
)

type Device struct {
	SDAPin  machine.Pin
	SCLPin  machine.Pin
	Address uint8
}

func New(SDAPin, SCLPin machine.Pin) *Device {
	return &Device{
		SDAPin:  SDAPin,
		SCLPin:  SCLPin,
		Address: ADDRESS,
	}
}

func (d *Device) delay() {
	time.Sleep(time.Microsecond * 2)
}

func (d *Device) Configure() {
	d.SDAPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.SCLPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	d.SDAPin.High()
	d.SCLPin.High()

	time.Sleep(time.Millisecond * 100)

	d.SendCommand(DISPLAYOFF)
	d.SendCommand(SETDISPLAYCLOCKDIV)
	d.SendCommand(0x80)
	d.SendCommand(SETMULTIPLEX)
	d.SendCommand(0x3F)
	d.SendCommand(SETDISPLAYOFFSET)
	d.SendCommand(0x0)
	d.SendCommand(SETSTARTLINE | 0x0)
	d.SendCommand(CHARGEPUMP)
	d.SendCommand(0x14)
	d.SendCommand(MEMORYMODE)
	d.SendCommand(0x00)
	d.SendCommand(SEGREMAP | 0x1)
	d.SendCommand(COMSCANDEC)
	d.SendCommand(SETCOMPINS)
	d.SendCommand(0x12)
	d.SendCommand(SETCONTRAST)
	d.SendCommand(0xCF)
	d.SendCommand(SETPRECHARGE)
	d.SendCommand(0xF1)
	d.SendCommand(SETVCOMDETECT)
	d.SendCommand(0x40)
	d.SendCommand(DISPLAYALLON_RESUME)
	d.SendCommand(NORMALDISPLAY)
	d.SendCommand(DISPLAYON)
}

func (d *Device) TxStart() {
	d.SDAPin.High()
	d.SCLPin.High()
	d.delay()
	d.SDAPin.Low()
	d.delay()
	d.SCLPin.Low()
	d.delay()
}

func (d *Device) TxStop() {
	d.SDAPin.Low()
	d.SCLPin.Low()
	d.delay()
	d.SCLPin.High()
	d.delay()
	d.SDAPin.High()
	d.delay()
}

func (d *Device) SendBit(bit bool) {
	if bit {
		d.SDAPin.High()
	} else {
		d.SDAPin.Low()
	}
	d.delay()

	d.SCLPin.High()
	d.delay()
	d.SCLPin.Low()
	d.delay()
}

func (d *Device) ReadACK() bool {
	d.SDAPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	d.delay()

	d.SCLPin.High()
	d.delay()
	ack := !d.SDAPin.Get()
	d.SCLPin.Low()
	d.delay()

	d.SDAPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return ack
}

func (d *Device) SendByte(data uint8) {
	for i := 7; i >= 0; i-- {
		bit := (data >> uint8(i)) & 0x01
		d.SendBit(bit == 1)
	}
	d.ReadACK()
}

func (d *Device) SendAddress(address uint8, write bool) {
	addr := address << 1
	if !write {
		addr |= 0x01
	}
	d.SendByte(addr)
}

func (d *Device) SendCommand(command uint8) {
	d.TxStart()
	d.SendAddress(d.Address, true)
	d.SendByte(0x00)
	d.SendByte(command)
	d.TxStop()
}

func (d *Device) DataStart() {
	d.TxStart()
	d.SendAddress(d.Address, true)
	d.SendByte(0x40)
}

func (d *Device) DataStop() {
	d.TxStop()
}

func (d *Device) DataByte(data uint8) {
	d.SendByte(data)
}

func (d *Device) SetPos(x uint8, y uint8) {
	d.SendCommand(COLUMNADDR)
	d.SendCommand(x)
	d.SendCommand(127)

	d.SendCommand(PAGEADDR)
	d.SendCommand(y / 8)
	d.SendCommand(7)
}

func (d *Device) SetPixel(x uint8, y uint8, color bool) {
	if x >= 128 || y >= 64 {
		return
	}
	d.SetPos(x, y)
	d.DataStart()
	if color {
		d.DataByte(0x80)
	} else {
		d.DataByte(0x00)
	}
	d.DataStop()
}

func (d *Device) FillScreen(fill_data uint8) {
	d.SetPos(0, 0)
	d.DataStart()
	for i := 0; i < 1024; i++ {
		d.DataByte(fill_data)
	}
	d.DataStop()
}

func (d *Device) DrawBuffer(x0 uint8, y0 uint8, x1 uint8, y1 uint8, bitmap []uint8) {
	if len(bitmap) == 0 {
		return
	}
	width := x1 - x0 + 1
	startPage := y0 / 8
	endPage := y1 / 8

	bitmapIndex := 0

	for page := startPage; page <= endPage; page++ {
		d.SendCommand(COLUMNADDR)
		d.SendCommand(x0)
		d.SendCommand(x1)

		d.SendCommand(PAGEADDR)
		d.SendCommand(page)
		d.SendCommand(page)

		d.DataStart()

		var data uint8
		for col := uint8(0); col < width; col++ {
			data = bitmap[bitmapIndex]
			bitmapIndex++
			d.DataByte(data)
		}
		d.DataStop()
	}
}

func (d *Device) Data(data uint8) {
	d.DataStart()
	d.DataByte(data)
	d.DataStop()
}

func (d *Device) ClearScreen() {
	d.FillScreen(0x00)
}

func (d *Device) Invert(invert bool) {
	if invert {
		d.SendCommand(INVERTDISPLAY)
	} else {
		d.SendCommand(NORMALDISPLAY)
	}
}

func (d *Device) SetContrast(contrast uint8) {
	d.SendCommand(SETCONTRAST)
	d.SendCommand(contrast)
}
