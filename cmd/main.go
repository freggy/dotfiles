package main

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/freggy/dotfiles/rootfs"
)

func main() {
	fs.WalkDir(rootfs.FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}
