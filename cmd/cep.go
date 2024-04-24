package cmd

import (
	"fast_api/app"
	// "fmt"

	"github.com/spf13/cobra"
)

var cepCmd = &cobra.Command{
	Use:   "cep",
	Short: "Inform your cep to search",
	Long:  "The cep passed will be used to call the APIs, and return its response",
	Run: func(cmd *cobra.Command, args []string) {
		// cep := args[0]
		// app.CallAPIs(cep)
		cep, err := cmd.Flags().GetString("command")
		if err != nil {
			panic(err)
		}

		app.CallAPIs(cep)
	},
}

func init() {
	rootCmd.AddCommand(cepCmd)
	cepCmd.Flags().StringP("command", "c", "", "Informe o cep")
	cepCmd.MarkFlagRequired("command")
}
