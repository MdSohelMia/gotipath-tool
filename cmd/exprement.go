/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/MdSohelMia/gotipath-tool/internal"
	"github.com/spf13/cobra"
)

// exprementCmd represents the exprement command
var exprementCmd = &cobra.Command{
	Use:   "exprement",
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
		internal.LargeFileCount(path)
	},
}

func init() {
	rootCmd.AddCommand(exprementCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exprementCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exprementCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
