package main

import (
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"

	"time"
)

type chord struct {
	name  string
	notes []midi.Note
}

var (
	buttonC  = machine.BUTTON_A
	buttonG  = machine.BUTTON_B
	buttonAm = machine.BUTTON_C
	buttonF  = machine.BUTTON_DOWN

	chords = []chord{
		{name: "C ", notes: []midi.Note{midi.C3, midi.E3, midi.G3}},
		{name: "G ", notes: []midi.Note{midi.G3, midi.B3, midi.D4}},
		{name: "Am", notes: []midi.Note{midi.A3, midi.C4, midi.E4}},
		{name: "F ", notes: []midi.Note{midi.F3, midi.A3, midi.C4}},
	}

	keys = []key{
		{name: "C", pin: buttonC, chord: chords[0]},
		{name: "G", pin: buttonG, chord: chords[1]},
		{name: "Am", pin: buttonAm, chord: chords[2]},
		{name: "F", pin: buttonF, chord: chords[3]},
	}

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	// change this to speed up or slow down the music
	bpm int = 80

	currentNote  int = 0
	currentNotes []midi.Note
)

func main() {
	usb.Product = "TinyGo Arpeggio"

	buttonC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	buttonG.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	buttonAm.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	buttonF.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	go handleDisplay()
	go handleKeys()

	for {
		playArpeggio()
		time.Sleep(time.Duration(60000/bpm/4) * time.Millisecond)
	}
}

func handleKeys() {
	for {
		for i := range keys {
			switch {
			case keys[i].press():
				startArpeggio(keys[i].chord.notes)

				keys[i].pressed = true

			case keys[i].release():
				stopArpeggio()

				keys[i].pressed = false

			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

type key struct {
	name    string
	pin     machine.Pin
	chord   chord
	pressed bool
}

func (k key) press() bool {
	if k.pin.Get() && !k.pressed {
		return true
	}
	return false
}

func (k key) release() bool {
	if !k.pin.Get() && k.pressed {
		return true
	}
	return false
}

func startArpeggio(notes []midi.Note) {
	if currentNotes != nil {
		stopArpeggio()
	}

	currentNotes = notes

	midi.Port().NoteOn(midicable, midichannel, currentNotes[currentNote], velocity)
}

func stopArpeggio() {
	if currentNotes == nil {
		return
	}

	midi.Port().NoteOff(midicable, midichannel, currentNotes[currentNote], velocity)

	currentNotes = nil
}

func playArpeggio() {
	if currentNotes == nil {
		return
	}

	midi.Port().NoteOff(midicable, midichannel, currentNotes[currentNote], velocity)

	currentNote++
	if currentNote >= len(currentNotes) {
		currentNote = 0
	}

	midi.Port().NoteOn(midicable, midichannel, currentNotes[currentNote], velocity)
}
