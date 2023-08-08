package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"machine"
	"machine/usb"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

//go:embed vial.json
var def []byte

func main() {
	usb.Product = "sgkb-0.4.0"

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	d := keyboard.New()

	colPins := []machine.Pin{
		machine.D0,
		machine.D1,
		machine.D2,
		machine.D3,
		machine.D4,
		machine.D5,
		machine.D8,
	}

	sm := d.AddSquaredMatrixKeyboard(colPins, [][]keyboard.Keycode{
		{
			jp.KeyEsc, jp.Key1, jp.Key2, jp.Key3, jp.Key4, jp.Key5, jp.Key6,
			jp.KeyTab, jp.KeyQ, jp.KeyW, jp.KeyE, jp.KeyR, jp.KeyT, 0,
			jp.KeyLeftCtrl, jp.KeyA, jp.KeyS, jp.KeyD, jp.KeyF, jp.KeyG, 0,
			jp.KeyLeftShift, jp.KeyZ, jp.KeyX, jp.KeyC, jp.KeyV, jp.KeyB, 0,
			jp.KeyMod1, jp.KeyLeftCtrl, jp.KeyWindows, jp.KeyLeftAlt, jp.KeyMod1, jp.KeySpace, 0,
			0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0,
		},
		{
			jp.KeyEsc, jp.KeyF1, jp.KeyF2, jp.KeyF3, jp.KeyF4, jp.KeyF5, jp.KeyF6,
			jp.KeyTab, jp.KeyQ, jp.KeyF15, jp.KeyEnd, jp.KeyF17, jp.KeyF18, 0,
			jp.KeyLeftCtrl, jp.KeyHome, jp.KeyS, jp.MouseRight, jp.MouseLeft, jp.MouseBack, 0,
			jp.KeyLeftShift, jp.KeyF13, jp.KeyF14, jp.MouseMiddle, jp.KeyF16, jp.MouseForward, 0,
			jp.KeyMod1, jp.KeyLeftCtrl, jp.KeyWindows, jp.KeyLeftAlt, jp.KeyMod1, jp.KeySpace, 0,
			0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0,
		},
	})
	sm.SetCallback(func(layer, index int, state keyboard.State) {
		fmt.Printf("sm: %d %d %d\n", layer, index, state)
	})

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{TX: machine.NoPin, RX: machine.UART_RX_PIN})

	uk := d.AddUartKeyboard(50, uart, [][]keyboard.Keycode{
		{
			0, jp.Key6, jp.Key7, jp.Key8, jp.Key9, jp.Key0, jp.KeyMinus, jp.KeyHat, jp.KeyBackslash2, jp.KeyBackspace,
			0, jp.KeyY, jp.KeyU, jp.KeyI, jp.KeyO, jp.KeyP, jp.KeyAt, jp.KeyLeftBrace, jp.KeyEnter, 0,
			0, jp.KeyH, jp.KeyJ, jp.KeyK, jp.KeyL, jp.KeySemicolon, jp.KeyColon, jp.KeyRightBrace, 0, 0,
			jp.KeyB, jp.KeyN, jp.KeyM, jp.KeyComma, jp.KeyPeriod, jp.KeySlash, jp.KeyBackslash, jp.KeyUp, jp.KeyDelete, 0,
			0, jp.KeySpace, jp.KeyHenkan, jp.KeyMod1, jp.KeyLeftAlt, jp.KeyPrintscreen, jp.KeyLeft, jp.KeyDown, jp.KeyRight, 0,
		},
		{
			0, jp.KeyF6, jp.KeyF7, jp.KeyF8, jp.KeyF9, jp.KeyF10, jp.KeyF11, jp.KeyF12, jp.KeyBackslash2, jp.KeyBackspace,
			0, jp.KeyY, jp.KeyU, jp.KeyTab, jp.KeyO, jp.WheelUp, jp.KeyAt, jp.KeyLeftBrace, jp.KeyEnter, 0,
			0, jp.KeyLeft, jp.KeyDown, jp.KeyUp, jp.KeyRight, jp.KeySemicolon, jp.KeyColon, jp.KeyRightBrace, 0, 0,
			jp.MouseForward, jp.WheelDown, jp.KeyM, jp.KeyComma, jp.KeyPeriod, jp.KeySlash, jp.KeyBackslash, jp.KeyPageUp, jp.KeyDelete, 0,
			0, jp.KeySpace, jp.KeyHenkan, jp.KeyMod1, jp.KeyLeftAlt, jp.KeyPrintscreen, jp.KeyHome, jp.KeyPageDown, jp.KeyEnd, 0,
		},
	})
	uk.SetCallback(func(layer, index int, state keyboard.State) {
		fmt.Printf("uk: %d %d %d\n", layer, index, state)
	})

	// override ctrl-h to BackSpace
	d.OverrideCtrlH()

	keyboard.KeyboardDef = []byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00, 0x00, 0x04, 0xE6, 0xD6, 0xB4, 0x46, 0x02, 0x00, 0x21, 0x01, 0x16, 0x00, 0x00, 0x00, 0x74, 0x2F, 0xE5, 0xA3, 0xE0, 0x02, 0xFB, 0x01, 0x26, 0x5D, 0x00, 0x3D, 0x88, 0x89, 0xC6, 0x54, 0x36, 0xC3, 0x17, 0x4F, 0xE4, 0xF9, 0xE8, 0x88, 0xA8, 0x34, 0x1C, 0xD8, 0xEE, 0x0A, 0x06, 0xA7, 0xF3, 0xD3, 0x76, 0xF3, 0x9E, 0xD8, 0x2A, 0x1D, 0xB3, 0x6E, 0xAD, 0x3A, 0x96, 0xB5, 0xAF, 0x10, 0x11, 0xFF, 0x76, 0x4B, 0x68, 0xD4, 0xED, 0xEB, 0x7C, 0x7C, 0xA7, 0x55, 0xB5, 0x36, 0x6A, 0x10, 0x87, 0x37, 0xA6, 0x5C, 0x00, 0x54, 0xB2, 0x86, 0x32, 0xB5, 0x2E, 0xF3, 0xE0, 0x3E, 0x8D, 0x88, 0x5A, 0x7D, 0x3F, 0x75, 0x6A, 0xA9, 0x52, 0x4C, 0x8E, 0x3D, 0x0E, 0xB0, 0x76, 0x22, 0x48, 0xE8, 0x62, 0x4A, 0x96, 0xEC, 0x2E, 0xCA, 0x53, 0x6C, 0xEE, 0x67, 0x3E, 0xC6, 0xE7, 0x75, 0xFB, 0xC9, 0xE5, 0x9D, 0xA2, 0x31, 0xF6, 0x53, 0x2A, 0x32, 0x6A, 0xFB, 0xFE, 0x0E, 0x65, 0x24, 0xC7, 0xC5, 0x52, 0x7F, 0xE2, 0x16, 0xF8, 0xFB, 0xAE, 0x58, 0xED, 0x6A, 0x7E, 0xC7, 0xC9, 0x54, 0xB1, 0xCB, 0x1F, 0x43, 0x7D, 0x23, 0x5B, 0xA2, 0x38, 0x6A, 0x23, 0xC6, 0x4E, 0xD8, 0x88, 0x4E, 0xDA, 0xA6, 0x75, 0x6B, 0x2C, 0xF9, 0x86, 0xAE, 0xE6, 0x01, 0x34, 0xC3, 0xC6, 0xAC, 0x1A, 0x87, 0x1A, 0x12, 0xB8, 0xAD, 0x13, 0x62, 0x54, 0x92, 0x61, 0x1C, 0x3E, 0x5D, 0x3C, 0x45, 0x1B, 0x28, 0xB4, 0xB0, 0x2A, 0x2B, 0x19, 0x1C, 0x59, 0x8E, 0xB1, 0x68, 0x70, 0x21, 0x96, 0xED, 0x70, 0xBB, 0xC4, 0x9B, 0xD7, 0x38, 0x4F, 0xC8, 0x2A, 0x68, 0xCF, 0xFD, 0x97, 0x71, 0x4B, 0xA7, 0x6C, 0xCD, 0xD6, 0x7A, 0x22, 0xD3, 0xB3, 0x0E, 0xC1, 0x9A, 0x00, 0xA3, 0x98, 0x1D, 0x85, 0x24, 0x47, 0xC4, 0x70, 0x3F, 0x59, 0x17, 0x42, 0xF0, 0x71, 0x7F, 0xEB, 0x28, 0xFF, 0x8D, 0x69, 0xC7, 0x15, 0xD1, 0x48, 0x36, 0x06, 0xEE, 0x8E, 0xD3, 0xB9, 0x3E, 0x02, 0x6A, 0x06, 0x48, 0xEF, 0xDE, 0x29, 0xA3, 0x7C, 0xF7, 0x6A, 0x37, 0x06, 0x1B, 0x98, 0x33, 0x9C, 0xF0, 0xB7, 0x60, 0xFB, 0xBE, 0xFE, 0x6B, 0xC2, 0xA0, 0xC3, 0xB2, 0xAE, 0x48, 0xF6, 0x2D, 0xC0, 0x61, 0xC0, 0xD1, 0x1D, 0x25, 0xF8, 0x1F, 0xED, 0x4B, 0x00, 0x00, 0x00, 0x00, 0xB3, 0xB3, 0x27, 0x1F, 0xF8, 0x65, 0x7D, 0x25, 0x00, 0x01, 0xC2, 0x02, 0xFC, 0x05, 0x00, 0x00, 0xF6, 0x53, 0xCE, 0x2F, 0xB1, 0xC4, 0x67, 0xFB, 0x02, 0x00, 0x00, 0x00, 0x00, 0x04, 0x59, 0x5A}

	return d.Loop(context.Background())
}
