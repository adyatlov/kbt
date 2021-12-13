package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kbt",
	Short: "Konvoy bundle tool",
	Long: "kbt is a CLI tool for analyzing Konvoy diagnostics bundle: " +
		"https://support.d2iq.com/hc/en-us/articles/4409215513236-Create-a-Konvoy-1-x-Support-Bundle",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
