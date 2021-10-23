package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "hydrator",
}

func init() {
	viper.SetConfigName("hydrator")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/hydrator/")
	viper.AddConfigPath("$HOME/.hydrator")
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("file not found ", err)
	} else {
		//fmt.Printf("%v\n", viper.AllSettings())
	}
}

func Execute() {
	rootCmd.AddCommand(getVersionCmd())
	rootCmd.AddCommand(getHydrateCmd())

	rootCmd.Execute()
}
