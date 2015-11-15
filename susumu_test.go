package susumu

import (
	"fmt"
	"time"
)

// This example shows a single progress bar rendering.
func ExampleBar_Draw() {
	b := &Bar{Width: 80, Max: 256, Position: 192, Unit: "KiB"}
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
	b := &Bar{Width: 80, Max: 256, Position: 0, Unit: "KiB"}
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
