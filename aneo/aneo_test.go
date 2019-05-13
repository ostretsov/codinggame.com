package aneo

import (
	"coding-game"
	"testing"
)

var testCases = []coding_game.TestCase{
	{
		"village traffic light",
		`50
1
200 15`,
		"50",
	},
	{
		"village traffic light 2",
		`50
1
200 10`,
		"36",
	},
	{
		"quiet country road",
		`90
3
300 30
1500 30
3000 30`,
		"90",
	},
	{
		"less quiet country road",
		`90
3
300 10
1500 10
3000 10`,
		"54",
	},
	{
		"rain of Traffic Lights",
		`130
100
500 15
1000 15
1500 15
2000 15
2500 15
3000 15
3500 15
4000 15
4500 15
5000 15
5500 15
6000 15
6500 15
7000 15
7500 15
8000 15
8500 15
9000 15
9500 15
10000 15
10500 15
11000 15
11500 15
12000 15
12500 15
13000 15
13500 15
14000 15
14500 15
15000 15
15500 15
16000 15
16500 15
17000 15
17500 15
18000 15
18500 15
19000 15
19500 15
20000 15
20500 15
21000 15
21500 15
22000 15
22500 15
23000 15
23500 15
24000 15
24500 15
25000 15
25500 15
26000 15
26500 15
27000 15
27500 15
28000 15
28500 15
29000 15
29500 15
30000 15
30500 15
31000 15
31500 15
32000 15
32500 15
33000 15
33500 15
34000 15
34500 15
35000 15
35500 15
36000 15
36500 15
37000 15
37500 15
38000 15
38500 15
39000 15
39500 15
40000 15
40500 15
41000 15
41500 15
42000 15
42500 15
43000 15
43500 15
44000 15
44500 15
45000 15
45500 15
46000 15
46500 15
47000 15
47500 15
48000 15
48500 15
49000 15
49500 15
50000 15`,
		"60",
	},
}

func TestSolve(t *testing.T) {
	for _, testCase := range testCases {
		coding_game.TestSolve(t, solve, testCase)
	}
}