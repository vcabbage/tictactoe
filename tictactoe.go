package tictactoe

import "fmt"

// Board is the state of the game
//
// Played positions are stored in the 9 right most of the 16 high order bits.
// The lower 16 bits is a mask of played positions.
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
	played := uint16(b >> 16) // extract high order bits

	switch player {
	case 'x', 'X':
	case 'o', 'O':
		played ^= 0xffff // if checking for 'O', invert the played positions
	default:
		panic(fmt.Sprintf("player must be X or O, not %q", player))
	}

	// apply the mask to get the positions user has played
	played &= uint16(b)

	// try each of the 8 possible winning orientations
	for _, check := range checks {
		if played&check == check {
			return true
		}
	}
	return false
}

// ArrayToBoard takes a 3x3 array of 'X's and 'O's
// and returns a Board. Any elements other than 'X'
// or 'O' are considered empty positions.
func ArrayToBoard(a [3][3]byte) Board {
	var b Board
	for x, row := range a {
		for y := range row {
			switch a[x][y] {
			case 'x', 'X':
				bit := uint(x*3 + y)
				b |= 1 << bit        // set position
				b |= 1 << (bit + 16) // set mask
			case 'o', 'O':
				b |= 1 << uint(x*3+y) // set mask
			}
		}
	}
	return b
}
