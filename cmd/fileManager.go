package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

func newFile(name ...string) {
	path := toPath(name...)
	dir, _ := filepath.Split(path)
	os.MkdirAll(dir, 0777)
	os.Create(path)
}

func newDir(name string) {
	os.MkdirAll(toPath(name), 0777)
}

func getMatching(pattern string) []string {
	files, _ := filepath.Glob(pattern)
	return files
}

func toPath(pathElmts ...string) string {
	path := filepath.Join(pathElmts...)
	if filepath.IsAbs(path) {
		return path
	}
	abs, _ := filepath.Abs(filepath.Join(homeDir, path))
	return abs
}

func toAbs(pathElmts ...string) string {
	abs, _ := filepath.Abs(filepath.Join(pathElmts...))
	return abs
}

func rmExt(file string) string {
	pieces := strings.Split(file, ".")
	return strings.Join(pieces[:len(pieces)-1], ".")
}
