package skynet

import (
	"coding-game"
	"testing"
)

func TestSolve(t *testing.T) {
	var testCases = []coding_game.TestCase{
		{
			"3 nodes, 1 gw",
			`3 2 1
0 1
1 2
2
0`,
			`2 1`,
		},
	}

	for _, tc := range testCases {
		coding_game.TestSolve(t, solve, tc)
	}
}
