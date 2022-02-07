// Copyright 2014 beego Author. All Rights Reserved.
//

package toolbox

import (
	"io"

	"github.com/W3-Engineers-Ltd/Radiant/core/admin"
)

// ProcessInput parse input command string
func ProcessInput(input string, w io.Writer) {
	admin.ProcessInput(input, w)
}

// MemProf record memory profile in pprof
func MemProf(w io.Writer) {
	admin.MemProf(w)
}

// GetCPUProfile start cpu profile monitor
func GetCPUProfile(w io.Writer) {
	admin.GetCPUProfile(w)
}

// PrintGCSummary print gc information to io.Writer
func PrintGCSummary(w io.Writer) {
	admin.PrintGCSummary(w)
}
