package utils

import "math"

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
