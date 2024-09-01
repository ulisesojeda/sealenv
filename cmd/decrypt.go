package cmd

import (
	"github.com/spf13/cobra"
	"sealenv/utils"
)

var env_file_decrypt string
var decryptCmd = &cobra.Command{
	Use:   "decrypt --env file",
	Short: "Print decrypted variables",
	Run: func(_cmd *cobra.Command, args []string) {
		password := utils.GetPassword()
		utils.UnsealVariables(env_file_decrypt, password)
	},
}

func init() {
	decryptCmd.Flags().StringVarP(&env_file_decrypt, "env", "", ".env", "environment file")

	rootCmd.AddCommand(decryptCmd)
}
