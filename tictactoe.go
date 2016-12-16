package tictactoe

import (
	"fmt"
	"sort"
)

// Board is the state of the game
//
// X positions are stored in the 16 high order bits,
// Y in the lower.
type Board uint32

var (
	// checkMasks contains masks for the 8 possible ways of winning
	checkMasks [8]uint16
	// checkMap is a map whose keys include every possible winning board
	checkMap map[uint16]struct{}
	// checkLookup is an array where with 2^9 elements, the value for indexes which
	// are a possible winning board is set to true
	checkLookup [512]bool
)

func init() {
	checkMasks = generateCheckMasks()
	checkMap = generateMap(checkMasks)
	checkLookup = generateLookup(checkMasks)
}

// IsWinner takes a Board and an 'X' or 'O' and returns whether
// the player has won.
//
// Implementation does a binary AND for each of the checks to determine
// if the board is a winner.
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
	for _, mask := range checkMasks {
		if played&mask == mask {
			return true
		}
	}
	return false
}

// IsWinnerMap is like IsWinner but checks for the board's existence
// in checkMap.
func IsWinnerMap(b Board, player byte) bool {
	switch player {
	case 'x', 'X':
		_, ok := checkMap[uint16(b>>16)]
		return ok
	case 'o', 'O':
		_, ok := checkMap[uint16(b)]
		return ok
	default:
		panic(fmt.Sprintf("player must be X or O, not %q", player))
	}
}

// IsWinnerLookup is like IsWinner but checks for the board's value
// is true in checkLookup.
func IsWinnerLookup(b Board, player byte) bool {
	switch player {
	case 'x', 'X':
		return checkLookup[uint16(b>>16)]
	case 'o', 'O':
		return checkLookup[uint16(b)]
	default:
		panic(fmt.Sprintf("player must be X or O, not %q", player))
	}
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

func generateMap(chks [8]uint16) map[uint16]struct{} {
	r := makePremutations(chks[:], 0)

	m := make(map[uint16]struct{}, 256)
	for _, chk := range r {
		m[chk] = struct{}{}
	}

	return m
}

func makePremutations(chks []uint16, n uint) []uint16 {
	if n == 8 {
		return chks
	}

	for _, chk := range chks {
		if (chk>>n)&1 != 1 {
			chks = append(chks, chk|1<<n)
		}
	}
	return makePremutations(chks, n+1)
}

type uint16s []uint16

func (a uint16s) Len() int           { return len(a) }
func (a uint16s) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint16s) Less(i, j int) bool { return a[i] < a[j] }

func generateLookup(chks [8]uint16) [512]bool {
	r := makePremutations(chks[:], 0)

	sort.Sort(uint16s(r))
	var s [512]bool // 2^9

	for _, chk := range r {
		s[chk] = true
	}
	return s
}
