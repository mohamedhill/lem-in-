package funcs

import (
	"bufio"
	"strings"
)

func ParseInput(scanner *bufio.Scanner) (*Graph, error) {
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasSuffix(line, "##") {
			continue
		}

	}
}
