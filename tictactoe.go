package tictactoe

import "fmt"

// Board is the state of the game
//
// X positions are stored in the 16 high order bits,
// Y in the lower.
type Board uint32

// checks holds winning masks, generated at initial load
var checks = generateCheckMasks()

func generateCheckMasks() [8]uint16 {
	chks := [8]uint16{
		6: uint16(0x0111), // 0000 0001 0001 0001 - diag down
		7: uint16(0x0054), // 0000 0000 0101 0100 - diag up
	}

	checkRow := uint16(0x01c0) // 0000 0001 1100 0000
	checkCol := uint16(0x0124) // 0000 0001 0010 0100
	for i := 0; i < 3; i++ {
		chks[i] = checkRow
		chks[i+3] = checkCol

		checkCol = checkCol >> 1
		checkRow = checkRow >> 3
	}

	return chks
}

// IsWinner takes a Board and an 'X' or 'O' and returns whether
// the player has won.
func IsWinner(b Board, player byte) bool {
	var played uint16

	switch player {
	case 'x', 'X':
		played = uint16(b >> 16) // extract high order bits
	case 'o', 'O':
		played = uint16(b)
	default:
		panic(fmt.Sprintf("player must be X or O, not %q", player))
	}

	// try each of the 8 possible winning orientations
	for _, check := range checks {
		if played&check == check {
			return true
		}
	}
	return false
}

// ArrayToBoard takes a 3x3 array of 'X's and 'O's
// and returns a Board.
func ArrayToBoard(a [3][3]byte) Board {
	var b Board
	for x, row := range a {
		for y := range row {
			switch a[x][y] {
			case 'x', 'X':
				b |= 1 << (uint(x*3+y) + 16)
			case 'o', 'O':
				b |= 1 << uint(x*3+y)
			}
		}
	}
	return b
}
