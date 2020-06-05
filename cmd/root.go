/*
Copyright © 2020 Mehdi Bounya <mehdi.bounya@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/mehdibo/go_ft/src/api"
	"github.com/mehdibo/go_ft/src/auth"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var outputFile string
var apiClient *api.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_ft",
	Short: "Bunch of tools to use the 42 API",
	Long: `With this tool you can easily fetch and send data from the 42 API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.go_ft.yaml)")
	rootCmd.PersistentFlags().StringVar(&outputFile, "file", "results.csv", "Where to store output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	requiredConfig := []string{
		"client_id",
		"client_secret",
		"scopes"}

	viper.SetDefault("token_endpoint", "https://api.intra.42.fr/oauth/token")
	viper.SetDefault("api_endpoint", "https://api.intra.42.fr/v2")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go_ft" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go_ft")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	// If a config file is found, read it in.
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error reading config file: \n%s\n", err)
		os.Exit(1)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	for _, config := range requiredConfig {
		if viper.Get(config) == nil {
			fmt.Println("'" + config + "' was not found in the config file but it is required.")
			os.Exit(1)
		}
	}

	apiClient = api.Create(viper.GetString("api_endpoint"), auth.GetOauthClient(viper.GetViper()))
}
