Given a Tic Tac Toe board and a player (X or O), returns whether the player has won.

Board is represented as a `uint32`. 9 bits are used as a mask to mark which positions have been played,
9 bits are used to represent whether an X or O was played. The remaining 14 bits are unused.

Download: `go get -v github.com/vcabbage/tictactoe`

Run tests: `go test -v github.com/vcabbage/tictactoe`
