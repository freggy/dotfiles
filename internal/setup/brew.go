package setup

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func InstallBrew(version string) error {
	if version == "" {
		version = "HEAD"
	}
	var (
		buf bytes.Buffer
		url = fmt.Sprintf("https://raw.githubusercontent.com/Homebrew/install/%s/install.sh", version)
	)
	if err := download(url, &buf); err != nil {
		return fmt.Errorf("download brew: %w", err)
	}
	if out, err := exec.Command("/bin/bash", buf.String()).CombinedOutput(); err != nil {
		log.Printf("error while installing brew: %s", string(out))
		return fmt.Errorf("install brew: %w", err)
	}
	return nil
}
