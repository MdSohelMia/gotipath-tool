/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/MdSohelMia/gotipath-tool/internal"
	"github.com/spf13/cobra"
)

// lcountCmd represents the lcount command
var lcountCmd = &cobra.Command{
	Use:   "lcount",
	Short: "A brief description of your command",
	Long:  `Large file`,
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
	rootCmd.AddCommand(lcountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lcountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lcountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
