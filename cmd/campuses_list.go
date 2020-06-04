/*
Copyright © 2020 Mehdi Bounya <mehdi.bounya@gmail.com>

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
	"github.com/gocarina/gocsv"
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
		// TODO: handle 429 too many requests
		if clientErr != nil {
			fmt.Fprintln(os.Stderr, clientErr)
			os.Exit(1)
		}

		respBody := helpers.GetRespBody(resp)

		var campuses []Campus

		err := json.Unmarshal(respBody, &campuses)
		if err != nil {
			fmt.Fprintf(os.Stderr, "There was an error parsing the JSON: \n%s\n", err)
			if len(campuses) == 0 {
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "But continuing anyway, the JSON was parsed")
		}

		content, err := gocsv.MarshalString(campuses)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't serialize to CSV: \n%s\n", err)
			os.Exit(1)
		}

		resultsFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0640)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't create file: \n%s\n", err)
			os.Exit(1)
		}

		_, err = resultsFile.WriteString(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't write data to file: \n%s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	campusesCmd.AddCommand(campusesListCmd)
}
