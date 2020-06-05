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
	"github.com/spf13/cobra"
	"time"
)

type Language struct {
	Id			uint `csv:"Lang_id"`
	Name		string `csv:"Lang_name"`
	Identifier	string `csv:"Lang_code"`
}

type Endpoint struct {
	Id uint `csv:"Endpoint_id"`
	Url string `csv:"Endpoint_url"`
	Description string `csv:"Endpoint_description"`
	CreatedAt *time.Time `csv:"Endpoint_created_at" json:"created_at"`
	UpdatedAt *time.Time `csv:"Endpoint_updated_at" json:"updated_at"`
}

type Campus struct {
	Id uint
	Name string
	TimeZone string `json:"time_zone"`
	Lang Language `json:"language"`
	UsersCount uint `json:"users_count"`
	VogsId uint `json:"vogsphere_id"`
	Country string
	Address string
	Zip string
	City string
	Website string
	Facebook string
	Twitter string
	Active bool
	Endpoint Endpoint
}

// campusesCmd represents the campuses command
var campusesCmd = &cobra.Command{
	Use:   "campuses",
	Short: "Interact with the campus entity",
}

func init() {
	rootCmd.AddCommand(campusesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campusesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campusesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
