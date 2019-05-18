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
			`2 1
`,
		},
		{
			"star",
			`11 23 1
0 1
0 2
0 10
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
9 10
10 1
11 1
11 2
11 3
11 4
11 5
11 6
11 7
11 8
11 9
11 10
11
0
10
9
8
`,
			`11 1
11 10
11 9
11 8
`,
		},
		{
			"star with 2 gateways",
			`11 23 2
0 1
0 2
0 10
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
9 10
10 1
11 1
11 2
11 3
11 4
11 5
11 6
11 7
11 8
11 9
11 10
3
8
0
10
9
11
`,
			`3 2
3 11
8 9
8 11
`,
		},
		{
			"two gateways, one isolated",
			`7 6 2
0 1
1 2
2 3
3 4
4 5
5 0
3
6
0`,
			`3 2
`,
		},
	}

	for _, tc := range testCases {
		coding_game.TestSolve(t, solve, tc)
	}
}
