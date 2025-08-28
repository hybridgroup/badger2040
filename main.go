package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/uc8151"
)

var display uc8151.Device
var btnA, btnB, btnC, btnUp, btnDown machine.Pin

var black = color.RGBA{1, 1, 1, 255}
var white = color.RGBA{0, 0, 0, 255}

const WIDTH = 296
const HEIGHT = 128

func main() {
	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	btnA = machine.BUTTON_A
	btnB = machine.BUTTON_B
	btnC = machine.BUTTON_C
	btnUp = machine.BUTTON_UP
	btnDown = machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	display.ClearBuffer()
	display.Display()
	setCustomData()

	tainigoLogo()
	time.Sleep(3 * time.Second)

	for {
		switch menu() {
		case 0:
			profile()
			break
		case 1:
			schedule(0, 0)
			break
		case 2:
			adventure()
			break
		case 3:
			demo()
			break
		case 4:
			wifiQR()
			break
		case 5:
			enterFlashingMode()
			break
		default:
			break
		}
		time.Sleep(1 * time.Second)
	}
}
