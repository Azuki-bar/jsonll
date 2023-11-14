package cli

import (
	"bufio"
	"encoding/json"
	"io"
)

func Main(stdin io.Reader, stdout, stderr io.Writer) error {
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		row := scanner.Bytes()
		if json.Valid(row) {
			stdout.Write(row)
      stdout.Write([]byte("\n"))
		}
	}
	return scanner.Err()
}
