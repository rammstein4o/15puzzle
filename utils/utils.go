package utils

import (
	"fmt"
	"math"
	"time"
)

func IsSolvable(seed []uint8) bool {
	puzzleLen := len(seed)
	puzzleCols := int(math.Sqrt(float64(puzzleLen)))
	parity := 0
	row := 0
	blankRow := 0

	for i, val := range seed {
		if i%puzzleCols == 0 {
			row++
		}
		if val == 0 {
			blankRow = row
			continue
		}

		for j := i + 1; j < puzzleLen; j++ {
			if val > seed[j] && seed[j] != 0 {
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

func GenerateSeed(size int) []uint8 {
	seed := []uint8{}
	for i := 0; i < size; i++ {
		seed = append(seed, uint8(i))
	}
	return seed
}
