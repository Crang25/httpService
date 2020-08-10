package router

import (
	"path"
	"strings"
)

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	index := strings.Index(p[1:], "/") + 1
	if index <= 0 {
		return p[1:], "/"
	}

	return p[1:index], p[index:]
}
