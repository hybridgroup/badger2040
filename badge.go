package main

import (
	"image/color"
	"math/rand"

	"github.com/conejoninja/badger2040/tetris"
	qrcode "github.com/skip2/go-qrcode"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinyfont/proggy"

	"tinygo.org/x/tinyfont/freesans"

	"tinygo.org/x/tinyfont/gophers"

	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

func profile() {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	badgeProfile()
	for {
		if btnA.Get() || btnB.Get() || btnC.Get() || btnUp.Get() || btnDown.Get() {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func demo() {
	badgeProfile()
	time.Sleep(30 * time.Second)
	blinky("Technologist", "for hire")
	myNameIs(YourName)
	time.Sleep(20 * time.Second)
	talkDate("AAA", "BBB", "CCC")
	sunrays()
	loading()
	tainigoLogo()
	time.Sleep(10 * time.Second)
	loadingInverted()
	tetrisfx()
	dvd("TinyGo")
}

func badgeProfile() {
	display.ClearBuffer()
	midW := int16(167)
	if profileErr == nil {
		img := pixel.NewImageFromBytes[pixel.Monochrome](128, 128, []byte(profileImg))
		if err := display.DrawBitmap(168, 0, img); err != nil {
			println(err)
		}
	}

	showRect(0, 0, midW, 30, black)
	tinydraw.Line(&display, 0, 0, 295, 0, black)
	tinydraw.Line(&display, 0, 0, 0, 127, black)
	tinydraw.Line(&display, 295, 0, 295, 127, black)
	tinydraw.Line(&display, midW, 0, midW, 127, black)
	tinydraw.Line(&display, 0, 87, midW, 87, black)
	tinydraw.Line(&display, 0, 107, midW, 107, black)
	tinydraw.Line(&display, 0, 127, 295, 127, black)

	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 8, 22, YourCompany, white)

	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, YourName)
	if w32 < uint32(midW) {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (midW-int16(w32))/2, 74, YourName, black)
	} else {
		w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, YourName)
		if w32 < uint32(midW) {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (midW-int16(w32))/2, 74, YourName, black)
		} else {
			w32, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, YourName)
			if w32 < uint32(midW) {
				tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (midW-int16(w32))/2, 74, YourName, black)
			} else {
				w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, YourName)
				tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (midW-int16(w32))/2, 74, YourName, black)
			}
		}
	}

	w32, _ = tinyfont.LineWidth(&freesans.Regular9pt7b, YourTitle)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, (midW-int16(w32))/2, 102, YourTitle, black)

	w32, _ = tinyfont.LineWidth(&freesans.Regular9pt7b, YourSocial)
	if w32 < uint32(midW) {
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, (midW-int16(w32))/2, 122, YourSocial, black)
	} else {
		w32, _ := tinyfont.LineWidth(&proggy.TinySZ8pt7b, YourSocial)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, (midW-int16(w32))/2, 120, YourSocial, black)
	}
	display.Display()
	display.WaitUntilIdle()
}

func showRect(x int16, y int16, w int16, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			display.SetPixel(i, j, c)
		}
	}
}

func blinky(topline, bottomline string) {
	display.ClearBuffer()

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, bottomline)
	for i := int16(0); i < 10; i++ {
		// fill the screen with with
		tinydraw.FilledRectangle(&display, 0, 0, WIDTH, HEIGHT, white)
		// show black text
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 50, topline, black)
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, black)

		// display
		display.Display()
		display.WaitUntilIdle()

		// repeat the other way around
		tinydraw.FilledRectangle(&display, 0, 0, WIDTH, HEIGHT, black)
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 50, topline, white)
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, white)

		display.Display()
		display.WaitUntilIdle()
	}
}

func myNameIs(name string) {
	display.ClearBuffer()

	var r int16 = 8

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, black)
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, black)
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, black)
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, black)

	// top band
	tinydraw.FilledRectangle(&display, r, 0, WIDTH-2*r-1, r, black)
	tinydraw.FilledRectangle(&display, 0, r, WIDTH, 26, black)

	// bottom band
	tinydraw.FilledRectangle(&display, r, HEIGHT-r-1, WIDTH-2*r-1, r+1, black)
	tinydraw.FilledRectangle(&display, 0, HEIGHT-2*r-1, WIDTH, r, black)

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&freesans.Regular12pt7b, "my NAME is")
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, int16(WIDTH-w32)/2, 24, "my NAME is", white)

	// middle text
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, name)
	if w32 < WIDTH {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 80, name, black)
	} else {
		w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, name)
		if w32 < WIDTH {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 80, name, black)
		} else {
			w32, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, name)
			if w32 < WIDTH {
				tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 80, name, black)
			} else {
				w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, name)
				tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 80, name, black)
			}
		}
	}

	tinyfont.WriteLine(&display, &gophers.Regular32pt, WIDTH-50, 110, "BE", black)

	display.Display()
}

func talkDate(topString, middleString, bottomString string) {
	display.ClearBuffer()

	// top text : Come see my talk on
	_, w32 := tinyfont.LineWidth(&freesans.Bold12pt7b, topString)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, int16(WIDTH-w32)/2, 28, topString, black)

	// middle text
	w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, middleString)
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, int16(WIDTH-w32)/2, 70, middleString, black)

	if bottomString != "" {
		// bottom text : at room XYZ
		w32, _ = tinyfont.LineWidth(&freesans.Regular12pt7b, bottomString)
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, int16(WIDTH-w32)/2, 110, bottomString, black)
	}

	display.Display()
	display.WaitUntilIdle()
}

func sunrays() {
	display.ClearBuffer()

	colors := [7][]color.RGBA{
		{white, white, white, white, white, white},
		{black, white, white, white, white, white},
		{black, white, white, black, white, white},
		{black, white, white, black, black, white},
		{black, black, white, black, black, white},
		{black, black, white, black, black, black},
		{black, black, black, black, black, black},
	}

	for i := 0; i < 21; i++ {
		if i%2 == 0 {
			tinydraw.FilledRectangle(&display, 0, 0, WIDTH, HEIGHT, white)
			rays(black)
		} else {
			tinydraw.FilledRectangle(&display, 0, 0, WIDTH, HEIGHT, black)
			rays(white)
		}
		tinydraw.FilledRectangle(&display, 20, 20, 256, 128, white)
		w32, _ := tinyfont.LineWidth(&freesans.Regular12pt7b, "Badge coded with")
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, int16(WIDTH-w32)/2, 50, "Badge coded with", black)

		w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, "TinyGo")
		tinyfont.WriteLineColors(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2, 100, "TinyGo", colors[i%7])

		display.Display()
		display.WaitUntilIdle()
	}

}

func rays(color color.RGBA) {
	// center point at the bottom
	var cx int16 = 148
	var cy int16 = 128

	// left side rays
	tinydraw.FilledTriangle(&display, cx, cy, 0, 0, 0, 22, color)
	tinydraw.FilledTriangle(&display, cx, cy, 0, 40, 0, 64, color)
	tinydraw.FilledTriangle(&display, cx, cy, 0, 80, 0, 108, color)

	// top rays
	tinydraw.FilledTriangle(&display, cx, cy, 22, 0, 44, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 66, 0, 88, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 110, 0, 132, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 154, 0, 176, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 198, 0, 220, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 242, 0, 264, 0, color)

	// right rays
	tinydraw.FilledTriangle(&display, cx, cy, 296, 10, 296, 32, color)
	tinydraw.FilledTriangle(&display, cx, cy, 296, 50, 296, 74, color)
	tinydraw.FilledTriangle(&display, cx, cy, 296, 90, 296, 118, color)
}

func loading() {
	display.ClearBuffer()

	for i := int16(0); i < 19; i++ {
		// draw a rectangle bigger each time
		tinydraw.FilledRectangle(&display, i*15, 0, 15, HEIGHT, black)

		// draw text again since a part of it was behind the rectangle
		w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "TinyGo")
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2, 70, "TinyGo", white)

		display.Display()
		display.WaitUntilIdle()
	}

	// LAST ONE IS SHORTER
	// draw a rectangle bigger each time
	tinydraw.FilledRectangle(&display, 285, 0, 11, HEIGHT, black)

	// draw text again since a part of it was behind the rectangle
	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "TinyGo")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2, 70, "TinyGo", white)

	display.Display()
	display.WaitUntilIdle()

}

func loadingInverted() {
	display.ClearBuffer()

	for i := int16(0); i < 19; i++ {
		// this is the opposite, we draw the text and draw a rectangle of the same color as the background
		// to make the ilusion the text is revealing
		w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "TinyGo")
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2, 70, "TinyGo", black)

		tinydraw.FilledRectangle(&display, i*15, 0, 250-i*15, HEIGHT, white)

		display.Display()
		display.WaitUntilIdle()
	}

	//THE LAS ONE IS DIFFERENT SIZE
	// this is the opposite, we draw the text and draw a rectangle of the same color as the background
	// to make the ilusion the text is revealing
	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "TinyGo")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2, 70, "TinyGo", black)

	tinydraw.FilledRectangle(&display, 11, 0, 11, HEIGHT, white)

	display.Display()
	display.WaitUntilIdle()
}

// tetrisfx
// will create a new ramdom piece ramdomly rotated each time the previous one stopped
func tetrisfx() {
	display.ClearBuffer()

	tetris.NewBoard()

	tetris.NewPiece()
	failed := 0
	k := 0
	for {
		display.ClearBuffer()
		if tetris.MovePiece() {
			failed = 0
		} else {
			failed++
			tetris.NewPiece()
		}
		tetris.DrawBoard(&display)
		tetris.DrawPiece(&display)

		w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "TinyGo")
		// add a white broder around the text
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2-2, 70-2, "TinyGo", white)
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2-2, 70+2, "TinyGo", white)
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2+2, 70-2, "TinyGo", white)
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2+2, 70+2, "TinyGo", white)
		// add the text
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, int16(WIDTH-w32)/2, 70, "TinyGo", black)

		display.Display()
		display.WaitUntilIdle()
		// stop after 5 pieces in a row that can not move (screen is kind of filled)
		if failed >= 5 {
			return
		}
		k++
	}
}

func dvd(text string) {
	display.ClearBuffer()

	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, text)
	maxW := int16(WIDTH - w32)
	maxH := int16(HEIGHT - 36) //assume line height is 36

	// random start point
	x := int16(rand.Int31n(int32(maxW)))
	y := int16(rand.Int31n(int32(maxH)))
	d := int16(4)
	dx := d
	dy := d

	for i := 0; i < 80; i++ { //duration 80 frames
		// paint white the previous text to "erase" it
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, x, y+36, "TinyGo", white)

		// move text
		x += dx
		y += dy

		// paint and show text
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, x, y+36, "TinyGo", black)
		display.Display()
		display.WaitUntilIdle()

		// change direction if needed
		if x >= maxW {
			dx = -d
		}
		if x <= 0 {
			dx = d
		}
		if y >= maxH {
			dy = -d
		}
		if y <= 0 {
			dy = d
		}
	}
}

func tainigoLogo() {
	display.ClearBuffer()

	img := pixel.NewImageFromBytes[pixel.Monochrome](246, 128, []byte(tainigo))
	if err := display.DrawBitmap(32, 0, img); err != nil {
		println(err)
	}

	display.Display()
	display.WaitUntilIdle()
}

func QR(msg string) {
	qr, err := qrcode.New(msg, qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(HEIGHT / len(qrbytes))

	bx := (WIDTH - size*factor) / 2
	by := (HEIGHT - size*factor) / 2
	display.ClearBuffer()
	for y := int16(0); y < size; y++ {
		for x := int16(0); x < size; x++ {
			if qrbytes[y][x] {
				tinydraw.FilledRectangle(&display, bx+x*factor, by+y*factor, factor, factor, black)
			} else {
				tinydraw.FilledRectangle(&display, bx+x*factor, by+y*factor, factor, factor, white)
			}
		}
	}
}
