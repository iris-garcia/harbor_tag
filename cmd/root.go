package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommand
var rootCmd = &cobra.Command{
	Use:   "harbor_tag",
	Short: "Retrieve tag information of Harbor registry artifacts",
	Long:  "Based on the current tags of the artifacts and the input from the user, generates the next tag",
}

// Execute - This is called by main.main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in ENV variables if set.
func initConfig() {
	// read in environment variables that match
	viper.AutomaticEnv()
}
