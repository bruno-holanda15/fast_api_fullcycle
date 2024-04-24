package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "fast_api",
	Short: "A caller to Cep APIs",
	Long: `The CLI calls for two APIs with a cep passed as an argument
	when executing the command, and returns the result from the fast API`,
}

func Execute() error{
	return rootCmd.Execute()
}