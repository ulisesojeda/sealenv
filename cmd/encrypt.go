package cmd

import (
	"github.com/spf13/cobra"
	"sealenv/utils"
)

var env_file_encrypt string
var out_file string
var encryptCmd = &cobra.Command{
	Use:   "encrypt --env file --out file",
	Short: "Generate file with encrypted variables",
	Run: func(_cmd *cobra.Command, args []string) {
		password := utils.GetPassword()
		utils.SealVariables(env_file_encrypt, out_file, password)
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&env_file_encrypt, "env", "", ".env", "environment file")
	encryptCmd.Flags().StringVarP(&out_file, "out", "", ".env.out", "encrypted environment file")
	rootCmd.AddCommand(encryptCmd)
}
