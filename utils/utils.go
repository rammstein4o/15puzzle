package utils

import (
	"fmt"
	"math"
	"time"
)

func IsSolvable(puzzle []uint8) bool {
	puzzleLen := len(puzzle)
	puzzleCols := int(math.Sqrt(float64(puzzleLen)))
	parity := 0
	row := 0
	blankRow := 0

	for i, val := range puzzle {
		if i%puzzleCols == 0 {
			row++
		}
		if val == 0 {
			blankRow = row
			continue
		}

		for j := i + 1; j < puzzleLen; j++ {
			if val > puzzle[j] && puzzle[j] != 0 {
				parity++
			}
		}
	}

	if puzzleCols%2 == 0 && blankRow%2 != 0 {
		return parity%2 != 0
	}

	return parity%2 == 0
}

func FormatDuration(d time.Duration) string {
	if d == 0 {
		return "00:00:00"
	}
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
