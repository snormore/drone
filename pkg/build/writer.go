package build

import (
	//"bytes"
	"fmt"
	"io"

	"strings"
)

var (
	// the prefix used to determine if this is
	// data that should be stripped from the output
	prefix = []byte("#DRONE:")
)

// custom writer to intercept the build
// output
type writer struct {
	io.Writer
}

// Write appends the contents of p to the buffer. It will
// scan for DRONE special formatting codes embedded in the
// output, and will alter the output accordingly.
func (w *writer) Write(p []byte) (n int, err error) {

	lines := strings.Split(string(p), "\n")
	for i, line := range lines {

		if strings.HasPrefix(line, "#DRONE:") {
			var cmd string

			// extract the command (base16 encoded)
			// from the output
			fmt.Sscanf(line[7:], "%x", &cmd)

			// echo the decoded command
			cmd = fmt.Sprintf("$ %s", cmd)
			w.Writer.Write([]byte(cmd))

		} else {
			w.Writer.Write([]byte(line))
		}

		if i < len(lines)-1 {
			w.Writer.Write([]byte("\n"))
		}
	}

	return len(p), nil
}

// WriteString appends the contents of s to the buffer.
func (w *writer) WriteString(s string) (n int, err error) {
	return w.Write([]byte(s))
}
