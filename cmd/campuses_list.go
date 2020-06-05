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
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/mehdibo/go_ft/src/helpers"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var campusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all campuses and write them to a results.csv files",
	Run: func(cmd *cobra.Command, args []string) {
		resp, clientErr := apiClient.Get("/campus")
		if clientErr != nil {
			fmt.Fprintln(os.Stderr, clientErr)
			os.Exit(1)
		}

		respBody := helpers.GetRespBody(resp)

		var campuses []Campus

		err := json.Unmarshal(respBody, &campuses)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "There was an error parsing the JSON: \n%s\n", err)
			os.Exit(1)
		}

		content, err := gocsv.MarshalString(campuses)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Couldn't serialize to CSV: \n%s\n", err)
			os.Exit(1)
		}

		helpers.WriteToFile(outputFile, content)
	},
}

func init() {
	campusesCmd.AddCommand(campusesListCmd)
}
