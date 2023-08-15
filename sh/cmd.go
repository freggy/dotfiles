package sh

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Cmd string

// Append multiple arguments to the existing command
// returning a completely new object.
// below is a basic example:
//
//	cat := sh.Command("cat")
//	cat.Exec(nil)
//	// now we want to cat a specific file
//	// but not copy the first statement again
//	cat.Append("/tmp/f.txt").Exec(nil)
func (c Cmd) Append(args ...string) Cmd {
	str := fmt.Sprintf("%s %s", string(c), strings.Join(args, " "))
	return Cmd(str)
}

// Exec runs the command with the given input.
// return values are identical to CombinedOutput()
func (c Cmd) Exec(in []byte) ([]byte, error) {
	parts := strings.Split(string(c), " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	if len(in) > 0 {
		cmd.Stdin = bytes.NewReader(in)
	}
	return cmd.CombinedOutput()
}
