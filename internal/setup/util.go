package setup

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func download(url string, buf *bytes.Buffer) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download file: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return fmt.Errorf("copy body: %w", err)
	}
	return nil
}
