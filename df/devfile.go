package df

import (
	"bufio"
	"fmt"
	"github.com/Sett17/dev/cli"
	"github.com/Sett17/dev/global"
	"os"
	"strings"
	"time"
)

func Read(path string) (ops []global.Operation, err error) {
	start := time.Now()
	file, err := os.Open(path)
	if err != nil {
		elapsed := time.Since(start)
		cli.Dbg(fmt.Sprintf("Parsing of %d operations took: %s", 0, elapsed))
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bufferOperation = global.Operation{}
	var scriptBuffer = strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "***") {
			if bufferOperation.Name != "" {
				bufferOperation.Script = scriptBuffer.String()
				ops = append(ops, bufferOperation)
				bufferOperation = global.Operation{}
				scriptBuffer.Reset()
			}
			trimmed := strings.TrimLeft(line, "*")
			nameDescription := trimmed[:strings.IndexRune(trimmed, '*')]
			bufferOperation.Name, bufferOperation.Description, _ = strings.Cut(nameDescription, "|")

			trimmed = strings.TrimLeft(strings.Replace(line, nameDescription, "", 1), "*")
			split := strings.Split(trimmed, "*")
			for _, s := range split {
				global.ParseOption(&bufferOperation, s)
			}
		} else {
			scriptBuffer.WriteString(line)
			scriptBuffer.WriteString("\n")
		}
	}
	bufferOperation.Script = scriptBuffer.String()
	ops = append(ops, bufferOperation)

	err = scanner.Err()
	elapsed := time.Since(start)
	cli.Dbg(fmt.Sprintf("Parsing of %d operations took: %s", len(ops), elapsed))
	return
}
