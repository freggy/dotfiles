package sh

import "os"

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
