package setup

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func InstallGo(version string) error {
	var (
		buf     bytes.Buffer
		tmpDest = "/tmp/go"
		url     = fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", version, runtime.GOOS, runtime.GOARCH)
	)
	if err := download(url, &buf); err != nil {
		return fmt.Errorf("download: %v", err)
	}
	if err := untar(&buf, tmpDest); err != nil {
		return fmt.Errorf("untar: %v", err)
	}
	if err := verify(version); err != nil {
		return fmt.Errorf("verify: %v", err)
	}
	if err := os.Rename(tmpDest, "/usr/local/go"); err != nil {
		return fmt.Errorf("rename: %w", err)
	}
	return nil
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
			// the go tarball does not contain TypeDir, but
			// we keep it just in case
			if err := os.MkdirAll(header.Name, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("mkdir: %w", err)
			}
		case tar.TypeReg:
			// the go tarball does not contain directories,
			// so we need to perform mkdir -p to ensure
			// all directories exist before creating the file
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
