package cmd

import (
	"fmt"
	"os"

	"github.com/adyatlov/kbt/csv"
	"github.com/spf13/cobra"
)

func podsToCSV(cmd *cobra.Command, _ []string) {
	var err error
	// Init reader
	r := os.Stdin
	i := cmd.Flag("input")
	if i.Changed {
		r, err = os.Open(i.Value.String())
	}
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Cannot open input file: %v", err.Error())
		os.Exit(1)
	}

	// Init writer
	w := os.Stdout
	o := cmd.Flag("output")
	O := cmd.Flag("default-name")
	if o.Changed && O.Changed {
		_, _ = fmt.Fprintln(os.Stderr, "ERROR: Flags -o (--output) and -O (--default-name) are mutually exclusive. "+
			"Please use only one of them.")
		os.Exit(1)
	}
	if o.Changed {
		w, err = os.Create(o.Value.String())
	}
	if O.Changed {
		w, err = os.Create("pods.csv")
	}
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Cannot create output file: %v", err.Error())
		os.Exit(1)
	}

	// Launching the command
	err = csv.PodsToCsv(r, w)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %v\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	var podsToCsv = &cobra.Command{
		Use:   "pods",
		Short: "Represents pods data as a CSV file",
		Long: "Prints information about pods in the CSV format. " +
			"If the output file is not specified, writes to the standard output",
		Run: podsToCSV,
	}
	podsToCsv.Flags().StringP("input", "i", "",
		"path to the pods YAML file. If path is not specified, reads from standard input.")
	podsToCsv.Flags().StringP("output", "o", "",
		"path to the output CSV file")
	podsToCsv.Flags().BoolP("default-name", "O", false,
		"write output to the pods.csv file")
	rootCmd.AddCommand(podsToCsv)
}
