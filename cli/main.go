/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/cloudybyte/shawty/cli/cmd"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Token string
}

func main() {
	cobra.OnInitialize(func() {
		dirname, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Couldn't obtain home directory")
			os.Exit(2)
		}
		viper.AddConfigPath(dirname + "/.config/shawty")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	})

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config File changed: ", e.Name)
	})

	cmd.Execute()
}
