package susumu_test

import (
	"fmt"
	"os"
	"time"

	"github.com/jsageryd/susumu"
)

// This example shows a single progress bar rendering.
func ExampleBar_Draw() {
	b := &susumu.Bar{Width: 80, Max: 256, Position: 192, Unit: "KiB", Output: os.Stdout}
	err := b.Draw()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	// Output:
	//  75% █████████████████████████████████████████████                192 / 256 KiB
}

// This is an example of an animated progress bar.
func ExampleBar_Draw_second() {
	b := &susumu.Bar{Width: 80, Max: 256, Position: 0, Unit: "KiB"}
	ticker := time.Tick(20 * time.Millisecond)
	for b.Position < b.Max {
		<-ticker
		b.Position++
		err := b.Draw()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println()
}
