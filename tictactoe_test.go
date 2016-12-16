package tictactoe_test

import (
	"testing"

	"github.com/vcabbage/tictactoe"
)

var tests = []struct {
	label   string
	board   [3][3]byte
	winnerX bool
	winnerO bool
}{
	{
		label: "O diagonal down",
		board: [3][3]byte{
			{'o', 0, 0},
			{0, 'o', 0},
			{0, 0, 'o'},
		},
		winnerO: true,
	},
	{
		label: "X diagonal up",
		board: [3][3]byte{
			{0, 0, 'x'},
			{0, 'x', 0},
			{'x', 0, 0},
		},
		winnerX: true,
	},
	{
		label: "O row 3",
		board: [3][3]byte{
			{0, 'x', 0},
			{'x', 0, 0},
			{'o', 'o', 'o'},
		},
		winnerO: true,
	},
	{
		label: "X row 1",
		board: [3][3]byte{
			{'x', 'x', 'x'},
			{0, 'x', 0},
			{'x', 0, 0},
		},
		winnerX: true,
	},
	{
		label: "X row 3",
		board: [3][3]byte{
			{0, 'x', 0},
			{'x', 0, 0},
			{'x', 'x', 'x'},
		},
		winnerX: true,
	},
	{
		label: "O col 2",
		board: [3][3]byte{
			{'o', 'o', 'x'},
			{0, 'o', 0},
			{'x', 'o', 0},
		},
		winnerO: true,
	},
	{
		label: "X col 3",
		board: [3][3]byte{
			{'o', 'o', 'x'},
			{0, 'x', 'x'},
			{'o', 'o', 'x'},
		},
		winnerX: true,
	},
	{
		label: "O col 1",
		board: [3][3]byte{
			{'o', 'o', 'x'},
			{'o', 'x', 'o'},
			{'o', 'o', 'x'},
		},
		winnerO: true,
	},
}

func TestIsWinner(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			board := tictactoe.ArrayToBoard(tt.board)

			if result := tictactoe.IsWinner(board, 'O'); tt.winnerO != result {
				t.Errorf("IsWinner(board, 'O') == %t, expected %t; board: %v\n%032b", result, tt.winnerO, tt.board, board)
			}

			if result := tictactoe.IsWinner(board, 'X'); tt.winnerX != result {
				t.Errorf("IsWinner(board, 'X') == %t, expected %t; board: %v\n%032b", result, tt.winnerX, tt.board, board)
			}
		})
	}
}
