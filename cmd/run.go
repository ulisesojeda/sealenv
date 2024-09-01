package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"sealenv/utils"
	"strings"
)

var env_file string
var runCmd = &cobra.Command{
	Use:   "run program --env file",
	Short: "Run program with env variables decrypted",
	Run: func(_cmd *cobra.Command, args []string) {
		var env [][]string
		var err error

		password := utils.GetPassword()

		raw, err := os.ReadFile(env_file)
		if err != nil {
			fmt.Println("Error reading environment file:", env_file)
			os.Exit(1)
		}
		data := string(raw)
		lines := strings.Split(data, "\n")

		for _, line := range lines {
			if len(line) > 0 {
				name_value, ok := utils.DecryptVariable(line, password)
				if ok {
					env = append(env, []string{name_value[0], name_value[1]})
				} else {
					fmt.Println("Invalid password")
					os.Exit(1)
				}
			}
		}

		cmd_list := strings.Split(args[0], " ")
		cmd := exec.Command(cmd_list[0], cmd_list[1:]...)

		cmd.Env = os.Environ()
		for _, variable := range env {
			s := fmt.Sprintf("%s=%s", variable[0], variable[1])
			cmd.Env = append(cmd.Env, s)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err = cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cmd.Wait()
	},
}

func init() {
	runCmd.Flags().StringVarP(&env_file, "env", "", ".env", "environment file")

	rootCmd.AddCommand(runCmd)
}
