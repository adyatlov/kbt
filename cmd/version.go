package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
	Date    string
)

func printVersion(*cobra.Command, []string) {
	fmt.Printf("kbt version: %v, commit: %v, release date: %v\n", Version, Commit, Date)
}

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show the Konvoy Bundle Tools version, commit, and release date",
		Run:   printVersion,
	}
	rootCmd.AddCommand(versionCmd)
}
