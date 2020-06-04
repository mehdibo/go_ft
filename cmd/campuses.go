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
	Id uint `json:"id"`
	Name string `json:"name"`
	TimeZone string `json:"time_zone"`
	Lang Language `json:"language"`
	UsersCount uint `json:"users_count"`
	VogsId uint `json:"vogsphere_id"`
	Country string `json:"country"`
	Address string `json:"address"`
	Zip string `json:"zip"`
	City string `json:"city"`
	Website string `json:"website"`
	Facebook string `json:"facebook"`
	Twitter string `json:"twitter"`
	Active bool `json:"active"`
	Endpoint Endpoint `json:"endpoint"`
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
