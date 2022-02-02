package utils

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	return utils.SelfPath()
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return utils.SelfDir()
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	return utils.FileExists(name)
}

// SearchFile Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	return utils.SearchFile(filename, paths...)
}

// GrepFile like command grep -E
// for example: GrepFile(`^hello`, "hello.txt")
// \n is striped while read
func GrepFile(patten string, filename string) (lines []string, err error) {
	return utils.GrepFile(patten, filename)
}
