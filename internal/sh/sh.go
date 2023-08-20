package sh

import (
	"os"
)

func Into(cmd Cmd, path string) error {
	out, err := cmd.Exec(nil)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0777)
}

func Exec(cmd Cmd) ([]byte, error) {
	return cmd.Exec(nil)
}

func ExecArgs(args ...string) ([]byte, error) {
	if len(args) <= 0 {
		return nil, nil
	}
	return Cmd(args[0]).Append(args[1:]...).Exec(nil)
}
