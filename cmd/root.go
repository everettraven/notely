package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "notely",
	Short: "A little note taking CLI",
}

func init() {
	rootCmd.AddCommand(&createCmd)
	rootCmd.AddCommand(&editCmd)
	rootCmd.AddCommand(&listCmd)
	rootCmd.AddCommand(&deleteCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
