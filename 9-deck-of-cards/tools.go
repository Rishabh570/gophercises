//go:build tools
// +build tools

// https://www.jvt.me/posts/2022/06/15/go-tools-dependency-management/
package main

import (
	_ "golang.org/x/tools/cmd/stringer"
)
