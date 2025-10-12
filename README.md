# Gopher ARCADE

![Gopher ARCADE](/assets/step13.jpg)

Gopher ARCADE is a simple ATTiny85 game console based on [Armour Grade](github.com/ArmourGrade/Attiny85-Arduino-Game-Console). I wanted to make some modifications and also improve my PCB making skills, so I recreated it from scratch since there were no source files, only gerbers.


There's a good collection of games written in Arduino/C: https://github.com/andyhighnumber/Attiny-Arduino-Games

My idea is to write some games in Go with [TinyGo](https://tinygo.org/)



## Building

While is pretty straight forward to mount all the components, check out the [build guide](BUILD_GUIDE.md) for the easiest order to do so.

## Hardware

In the _hardware/kicad_ folder you could find the KiCAD project with the used mods as well as the gerbers files in case you want to just order some PCBs without modifying them.


### Parts list

| No | Item name                        | Quantity | Notes |
|----|----------------------------------|------|------|
| 1  | PCB                              | 1    | 
| 2  | ATTiny85 20PU                    | 1    | 
| 3  | DIP8P socket                     | 1    | 
| 4  | SSD1306 I²C 128x64 display       | 1    | 
| 5  | 10k ohm resistors (through hole) | 2    | 
| 6  | SN74LVC2G04DBVR                  | 1    | 
| 7  | CPT-1203-78-SMT-TR (buzzer)      | 1    | 
| 8  | EG1213 switch                    | 1    | 
| 9  | CR2032 battery holder (smd)      | 1    | 
| 10 | Kailh choc v2 low profile        | 2    | footprint accepts both kailh choc v1 & v2
| 11 | Kailh choc v2 low profile keycap | 2    | 
| 12 | Display spacer (3d printed)      | 1    | optional, but helps aligning the display

## Software

### Code

The _github.com/conejoninja/arcade/tinyssd1306_ package has a simple driver for the display as the once in the [drivers repository](https://github.com/tinygo-org/drivers) has a buffer and doesn't fit in the **ATTiny85**.

Board's pins are defined as follows:

```go
const (
	P5 Pin = PB0
	P6 Pin = PB1
	P7 Pin = PB2
	P2 Pin = PB3
	P3 Pin = PB4
	P1 Pin = PB5

	LED          = P1
	BUTTON_LEFT  = P7
	BUTTON_RIGHT = P5
	SPEAKER      = P6
)
```

### Flashing


## Requirements

By default the ATTiny85 runs at 1MHz, you can make it runs at 8MHz with the following command (this is only required once).

```bash
avrdude -p attiny85 -c usbasp -B 32 -U lfuse:w:0xE2:m
```


You could use [TinyGo](https://tinygo.org/) with USBASP and the target _gopher-arcade_ to flash it as follows:

```bash
tinygo flash -size=short -target gopher-arcade  .
```

