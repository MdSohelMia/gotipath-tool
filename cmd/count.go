/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/MdSohelMia/gotipath-tool/internal"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string

		if len(args) == 1 {
			path = args[0]
		} else {
			p, _ := os.Getwd()
			path = p
		}
		internal.Walk(path)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
