/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"

	"github.com/cloudybyte/shawty/cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Pretty much does what it says",
	Run: func(cmd *cobra.Command, args []string) {
		c := api.NewClient(&http.Client{})
		if len(args) < 2 {
			cmd.PrintErr("Must provide username and password")
			return
		} else if len(args) > 2 {
			cmd.PrintErr("Must provide only username and password")
			return
		}

		sess, err := c.Login(args[0], args[1])
		if err != nil {
			cmd.PrintErr("Couldn't create session: ", err)
		}
		viper.GetViper().Set("token", sess.Token)
		//TODO: we probably shouldn't write the session token in plaintext to disk. Maybe use the keyring here?
		if err := viper.WriteConfig(); err != nil {
			cmd.PrintErr("Couldn't write config: ", err)
		}
		cmd.Printf("Logged in as: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
