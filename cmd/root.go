package cmd

import (
	"fmt"
	"my-cli/github"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli",
	Short: "Uma aplicação CLI para buscar usuários no GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		if user == "" {
			fmt.Println("É necessário informar um usuário com -u ou --user")
			os.Exit(1)
		}
		github.GetUser(user)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("user", "u", "", "Usuário do GitHub")
}

func initConfig() {
	viper.AutomaticEnv()
}
