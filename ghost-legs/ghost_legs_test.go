package ghost_legs

import (
	"coding-game"
	"testing"
)

var testCases = []coding_game.TestCase{
	{
		"no links",
		`7 7
A  B  C
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
1  2  3`,
		`A1
B2
C3
`,
	},
	{
		"example input",
		`7 7
A  B  C
|  |  |
|--|  |
|  |--|
|  |--|
|  |  |
1  2  3`,
		`A2
B1
C3
`,
	},
	{
		"a little bit harder",
		`7 7
A  B  C
|--|  |
|  |--|
|--|  |
|  |--|
|--|  |
1  2  3`,
		`A1
B3
C2
`,
	},
	{
		"small sample",
		`13 8
A  B  C  D  E
|  |  |  |  |
|  |--|  |  |
|--|  |  |  |
|  |  |--|  |
|  |--|  |--|
|  |  |  |  |
1  2  3  4  5`,
		`A3
B5
C1
D2
E4
`,
	},
}

func TestSolve(t *testing.T) {
	for _, testCase := range testCases {
		coding_game.TestSolve(t, solve, testCase)
	}
}
