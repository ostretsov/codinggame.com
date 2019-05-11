package ghost_legs

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

var (
	errNoLinkedLine             = errors.New("there is no link")
	diagramWidth, diagramHeight int
	lines                       = []byte{}
	lineEndings                 = map[byte]byte{}
	links                       = []link{}
)

type link struct {
	from, to byte
}

func (l *link) getLinkedLine(x byte) (byte, error) {
	if x == l.from {
		return l.to, nil
	}
	if x == l.to {
		return l.from, nil
	}
	return byte(0), errNoLinkedLine
}

// 3. Why if with * caller would be incorrect
func solve(in io.Reader, out io.Writer) {
	resetVars()
	parseInput(in)

	for _, line := range lines {
		currentLine := line
		for _, l := range links {
			linkedLine, err := l.getLinkedLine(currentLine)
			if err != nil {
				continue
			}
			currentLine = linkedLine
		}
		fmt.Fprintf(out, "%s%s\n", string(line), string(lineEndings[currentLine]))
	}
}

func resetVars() {
	lines = []byte{}
	lineEndings = map[byte]byte{}
	links = []link{}
}

func parseInput(in io.Reader) {
	fmt.Fscanf(in, "%d %d\n", &diagramWidth, &diagramHeight)

	// parse line names
	l := make([]byte, diagramWidth+1) // includes new line
	in.Read(l)
	for i := 0; i < diagramWidth; i += 3 {
		lines = append(lines, l[i])
	}

	for i := 0; i < diagramHeight-2; i++ {
		in.Read(l)
		for j, lineIndex := 1, 0; j < diagramWidth; j, lineIndex = j+3, lineIndex+1 {
			if l[j] == '-' {
				links = append(links, link{lines[lineIndex], lines[lineIndex+1]})
			}
		}
	}

	// parse line endings
	in.Read(l)
	for i, lineIndex := 0, 0; i < diagramWidth; i, lineIndex = i+3, lineIndex+1 {
		lineEndings[lines[lineIndex]] = l[i]
	}
}
