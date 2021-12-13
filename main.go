package main

import (
	"github.com/adyatlov/kbt/cmd"
)

// GoReleaser populates these vars.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date
	cmd.CheckNewRelease()
	cmd.Execute()
}
