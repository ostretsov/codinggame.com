package coding_game

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type TestCase struct {
	Name            string
	In, ExpectedOut string
}

func TestSolve(t *testing.T, solveFunc func(reader io.Reader, writer io.Writer), testCase TestCase) {
	t.Run(testCase.Name, func(t *testing.T) {
		in := strings.NewReader(testCase.In)
		out := bytes.NewBuffer([]byte{})

		solveFunc(in, out)

		got := string(out.Bytes())
		expected := testCase.ExpectedOut
		if strings.Compare(expected, got) != 0 {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
