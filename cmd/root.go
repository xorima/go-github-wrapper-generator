/*
Copyright © 2023 Jason Field <xorima@xorima.dev>
*/
package cmd

import (
	"os"

	"github.com/xorima/go-github-wrapper-generator/generator"

	"github.com/spf13/cobra"
)

var version string
var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-github-wrapper-generator",
	Short: "A tool to help regenerate the go-github-wrapper repository",
	Long: `This tool will generate a structed based wrapper for the go-github module
	this is to allow easier interfacing of the go-github module with other projects and thus enable 
	a simpler mocking approach for testing.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		h := generator.NewGenerator(path, version)
		h.Handle()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&version, "version", "v", "v53", "The version of the github package to use")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "The path on disk to put the generated files")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
