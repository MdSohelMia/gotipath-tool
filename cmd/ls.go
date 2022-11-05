/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/MdSohelMia/gotipath-tool/internal"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List the objects in the path with size and path.",
	Long: `Lists the objects in the source path to standard output in a human
	readable format with size and path.`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string

		if len(args) == 1 {
			path = args[0]
		} else {
			p, _ := os.Getwd()
			path = p
		}
		internal.Ls(path)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
