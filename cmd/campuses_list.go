/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mehdibo/go_ft/src/auth"
	"github.com/mehdibo/go_ft/src/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// listCmd represents the list command
var campusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all campuses and write them to a results.csv files",
	Run: func(cmd *cobra.Command, args []string) {
		// Create an authenticated http client
		client := auth.GetOauthClient(viper.GetViper())

		resp, clientErr := client.Get(viper.GetString("api_endpoint")+"/campus")
		if clientErr != nil {
			fmt.Fprintln(os.Stderr, clientErr)
			os.Exit(1)
		}

		/**
		 * TODO: better organize serializers and normalizers to add possibility for other formats and options
		 */
		respBody := helpers.GetRespBody(resp)

		var campuses []Campus

		json.Unmarshal(respBody, &campuses)

		// Print keys
		var keys []string
		for key, _ := range CampusNormalizer(campuses[0]) {
			keys = append(keys, key)
		}
		helpers.CsvWriter(keys)


		for _, campus := range campuses {
			var line []string
			normalizedCampus := CampusNormalizer(campus)
			for _, key := range keys {
				line = append(line, normalizedCampus[key])
			}
			helpers.CsvWriter(line)
		}
	},
}

func init() {
	campusesCmd.AddCommand(campusesListCmd)
}