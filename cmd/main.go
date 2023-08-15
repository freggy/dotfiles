package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	// TODO: cleanup tmp after every fail
	version := "1.21.1"

	var (
		buf     bytes.Buffer
		tmpDest = "/tmp/go"
		url     = fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", version, runtime.GOOS, runtime.GOARCH)
	)

	if err := download(url, &buf); err != nil {
		log.Fatalf("download: %v", err)
	}

	if err := untar(&buf, tmpDest); err != nil {
		log.Fatalf("untar: %v", err)
	}

	if err := verify(version); err != nil {
		log.Fatalf("verify: %v", err)
	}

	if err := os.Rename(tmpDest, "/usr/local/go"); err != nil {
		log.Fatalf("rename: %w", err)
	}
}

func untar(buf *bytes.Buffer, dest string) error {
	if err := os.MkdirAll(dest, 0777); err != nil {
		return fmt.Errorf("mkdir dest: %w", err)
	}

	gzr, err := gzip.NewReader(buf)
	if err != nil {
		return fmt.Errorf("gzip: %w", err)
	}
	r := tar.NewReader(gzr)

	for {
		header, err := r.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return fmt.Errorf("read tar: %w", err)
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(header.Name, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("mkdir: %w", err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(header.Name), 0777); err != nil {
				return fmt.Errorf("mkdir open file: %w", err)
			}

			f, err := os.OpenFile(header.Name, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("open file: %w", err)
			}

			if _, err := io.Copy(f, r); err != nil {
				return fmt.Errorf("copy: %w", err)
			}

			f.Close()
		}
	}
}

// verify verifies that the tarball we installed
// actually contains the binary we expect.
func verify(version string) error {
	out, err := exec.Command("./go/bin/go", "version").CombinedOutput()
	if err != nil {
		return err
	}
	expected := fmt.Sprintf("go version go%s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	if string(out) != expected {
		return fmt.Errorf("unexpected version string, expected: '%s' got '%s'", expected, out)
	}
	return nil
}

func download(url string, buf *bytes.Buffer) (ret error) {
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
