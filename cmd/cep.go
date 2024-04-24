package cmd

import (
	"fast_api/app"

	"github.com/spf13/cobra"
)

var cepCmd = &cobra.Command{
	Use:   "cep",
	Short: "Inform your cep to search",
	Long:  "The cep passed will be used to call the APIs, and return its response",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cep := args[0]
		app.CallAPIs(cep)
	},
}

func init() {
	rootCmd.AddCommand(cepCmd)
}
