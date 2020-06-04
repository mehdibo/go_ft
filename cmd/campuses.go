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
	"fmt"
	"github.com/spf13/cobra"
	"reflect"
)

type Language struct {
	Id			uint
	Name		string
	Identifier	string
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
	Endpoint string `json:"endpoint"`
}

// campusesCmd represents the campuses command
var campusesCmd = &cobra.Command{
	Use:   "campuses",
	Short: "Interact with the campus entity",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("campuses called")
	//},
}

func CampusNormalizer(campus Campus) map[string]string {
	var normalizedCampus = make(map[string]string)

	fields := reflect.Indirect(reflect.ValueOf(campus))

	for i:= 0; i < fields.NumField(); i++ {
		var key = fields.Type().Field(i).Name
		var val string
		fieldType := fields.Field(i).Type().Name()
		switch fieldType {
			case "uint":
				val = fmt.Sprintf("%d", fields.Field(i).Interface())
			case "bool":
				val = fmt.Sprintf("%t", fields.Field(i).Interface())
			case "Language":
				continue
			default:
				val = fmt.Sprintf("%s", fields.Field(i).Interface())
		}
		normalizedCampus[key] = val
	}
	// TODO: normalize language

	return normalizedCampus
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
