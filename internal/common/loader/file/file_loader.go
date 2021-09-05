package file

import (
	"bufio"
	"os"

	"github.com/lexcao/watch-log/pkg/component"
)

type LiveTailLoader struct {
}

type scannerBased struct {
	scanner *bufio.Scanner
}

func (s scannerBased) HasNext() bool {
	return s.scanner.Scan()
}

func (s scannerBased) Next() string {
	return s.scanner.Text()
}

func (l LiveTailLoader) Load(file *os.File) component.IteratorLine {
	// TODO handle live tail from file
	// First support pipe
	scanner := bufio.NewScanner(bufio.NewReader(file))

	return scannerBased{scanner}
}
