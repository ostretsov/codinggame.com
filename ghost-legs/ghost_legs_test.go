package ghost_legs

import (
	"bytes"
	"strings"
	"testing"
)

type testCase struct {
	name            string
	in, expectedOut string
}

var testCases = []testCase{
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
		t.Run(testCase.name, func(t *testing.T) {
			in := strings.NewReader(testCase.in)
			out := bytes.NewBuffer([]byte{})

			solve(in, out)

			got := string(out.Bytes())
			expected := testCase.expectedOut
			if strings.Compare(expected, got) != 0 {
				t.Errorf("expected %v, got %v", expected, got)
			}
		})
	}
}
