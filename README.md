Given a Tic Tac Toe board and a player (X or O), returns whether the player has won.

Board is represented as a `uint32`. 9 bits are used to mark X positions,
and 9 bits are used to represent O positions. The remaining 14 bits are unused.

Three approaches are included:

* `IsWinner`: Checks the user's positions against 8 masks representing the 8 possible ways to win.
* `IsWinnerMap`: Checks for the user's positions against a pre-computed map whose keys include all possible winning positions.
* `IsWinnerLookup`; Checks for the user's positions against a precomputed array of all positions, with winning positions holding a value of true.

Download: `go get -v github.com/vcabbage/tictactoe`

Run tests: `go test -v github.com/vcabbage/tictactoe`

Run benchmarks: `go test -v -run=xxx -bench=. github.com/vcabbage/tictactoe`

Benchmarks from 2015 MacBook Pro:
```
BenchmarkIsWinner-8         	100000000	        74.8 ns/op
BenchmarkIsWinnerMap-8      	50000000	       120 ns/op
BenchmarkIsWinnerLookup-8   	100000000	        68.5 ns/op
```