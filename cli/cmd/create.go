/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/cloudybyte/shawty/cli/api"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new shortlink",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := api.NewClient(&http.Client{})
		if len(args) < 1 {
			cmd.PrintErr("Must provide target")
			return
		} else if len(args) > 1 {
			cmd.PrintErr("Must only provide target")
			return
		}
		red, err := c.CreateRedirect(args[0])
		if err != nil {
			switch err.Error() {
			case "unauthorized":
				fmt.Printf("Status: %s\nUnauthorized. Please try logging back in.\n", "401")
				return
			case "internal_server_error":
				fmt.Printf("Status: %s\nThis means sth is wrong on the servers end. You are most likely not at fault here and cannot fix the error.\nPlease contact your admin.\n", "500")
				return
			default:
				fmt.Println("An error occured. Please check the stacktrace and contact the developer.")

				panic(err)
			}
		}
		cmd.Printf("Successfully shortened %s to %s\n", args[0], red.Short)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
