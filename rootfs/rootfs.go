package rootfs

import (
	"embed"
	_ "embed"
)

// FS holds all files of the filesystem tree in this folder.
// see why we need `all:` here:
// * https://go-review.googlesource.com/c/go/+/359413
//
//go:embed all:*
var FS embed.FS
