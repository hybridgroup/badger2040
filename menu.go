package main

import (
	"time"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

func menu() int16 {
	options := []string{
		"Badge",
		"GopherCon Schedule",
		"GopherCon Adventure",
		"Demo",
		"Wifi QR",
		"Flashing mode (UF2)",
	}

	display.ClearBuffer()

	showRect(0, 0, WIDTH, 16, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 6, 12, "MENU", white)

	selected := int16(0)
	numOpts := int16(len(options))
	for i := int16(0); i < numOpts; i++ {
		tinydraw.Circle(&display, 20, 38+12*i, 3, black)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 28, 40+12*i, options[i], black)
	}
	tinydraw.FilledCircle(&display, 20, 38, 4, black)

	img := pixel.NewImageFromBytes[pixel.Monochrome](112, 96, []byte(logochip))
	if err := display.DrawBitmap(168, 24, img); err != nil {
		println(err)
	}

	display.Display()
	display.WaitUntilIdle()

	released := true
	for {
		if released && btnUp.Get() && selected > 0 {
			selected--
			tinydraw.FilledCircle(&display, 20, 38+12*selected, 4, black)
			tinydraw.FilledCircle(&display, 20, 38+12*(selected+1), 4, white)
			tinydraw.Circle(&display, 20, 38+12*(selected+1), 3, black)

			display.Display()
		}
		if released && btnDown.Get() && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(&display, 20, 38+12*selected, 4, black)
			tinydraw.FilledCircle(&display, 20, 38+12*(selected-1), 4, white)
			tinydraw.Circle(&display, 20, 38+12*(selected-1), 3, black)

			display.Display()
		}
		if released && (btnA.Get() || btnB.Get() || btnC.Get()) {
			break
		}
		if !btnA.Get() && !btnB.Get() && !btnC.Get() && !btnUp.Get() && !btnDown.Get() {
			released = true
		} else {
			released = false
		}
		time.Sleep(200 * time.Millisecond)
	}
	return selected
}
