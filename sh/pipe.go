package sh

import "os"

type Pipe struct {
	Err error
	Out []byte
}

func P() *Pipe {
	return &Pipe{}
}

// P functions the same as the pipe (|) operator
// in the shell.
// below is a basic example:
//
//	// echo hello | cat > /tmp/file
//	sh.P().P("echo hello").P("cat").Into("/tmp/file")
//
//	// echo hello | cat
//	sh.P().P("echo hello").P("cat")
func (p *Pipe) P(cmd Cmd) *Pipe {
	// this basically implements set -o pipefail behavior
	// we have to guard every pipe function with this
	if p.Err != nil {
		return p
	}
	p.Out, p.Err = cmd.Exec(p.Out)
	return p
}

// Into functions the same as the (>) redirect operator.
// below is a basic example:
//
//	// echo hello | cat > /tmp/file
//	sh.P().P("echo hello").P("cat").Into("/tmp/file")
func (p *Pipe) Into(path string) error {
	if p.Err != nil {
		return p.Err
	}
	return os.WriteFile(path, p.Out, 0777)
}
