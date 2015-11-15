// Package susumu implements a simple progress bar.
package susumu

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Bar is a progress bar.
type Bar struct {
	Width    int
	Max      int
	Position int
	Unit     string
}

var blocks = [9]rune{' ', '▏', '▎', '▍', '▌', '▋', '▊', '▉', '█'}

// Draw draws the progress bar. Update Position then redraw as needed.
func (b *Bar) Draw() error {
	widthOfMax := int(math.Log10(float64(b.Max))) + 1
	barWidth := b.Width - 11 - widthOfMax*2 - len(b.Unit)
	if barWidth <= 0 {
		return errors.New("bar does not fit within given width")
	}
	totalSteps := barWidth * 8
	totalBlockCount := totalSteps / 8
	posStringFormat := fmt.Sprintf("%%%dd", widthOfMax)

	translatedPosition := (totalSteps * b.Position) / b.Max
	percent := (100 * translatedPosition) / totalSteps
	blockCount := translatedPosition / 8
	lastBlockSize := int(math.Mod(float64(translatedPosition), 8))

	bar := strings.Repeat(string(blocks[8]), blockCount) + string(blocks[lastBlockSize])
	afterPadding := strings.Repeat(" ", totalBlockCount-blockCount)
	fmt.Printf(" %3d%% %s%s"+posStringFormat+" / %d %s\r", percent, bar, afterPadding, b.Position, b.Max, b.Unit)
	return nil
}
