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
	"github.com/go-playground/validator/v10"
	"github.com/mehdibo/go_ft/src/helpers"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
)

type newUser struct {
	Email string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Kind string `json:"kind" validate:"oneof=admin student external"`
	CampusId uint64 `json:"campus_id"`
}

type createUserReq struct {
	User newUser  `json:"user"`
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create email first_name last_name type campus",
	Short: "Create a new user account",
	Long: `Create a new user account
type: must be either admin, student or external
campus: the campus id where the user belongs`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		campusId, err := strconv.ParseUint(args[4], 10, 0)
		if err != nil {
			helpers.PrintfErrorExit("campus argument must be a valid integer and greater than 0")
		}
		reqBody := createUserReq{
			User: newUser{
				Email: args[0],
				FirstName: args[1],
				LastName: args[2],
				Kind: args[3],
				CampusId: campusId,
			},
		}
		v := validator.New()

		validationErr := v.Struct(reqBody)
		if validationErr != nil {
			helpers.PrintfErrorExit("%s", validationErr)
		}

		resp, err := apiClient.PostJson("/users", reqBody)
		if err != nil {
			helpers.PrintfErrorExit("%s", err)
		}

		switch resp.StatusCode {
			case http.StatusOK:
				fmt.Println("The user was created successfully!")
				os.Exit(0)
			case http.StatusForbidden:
				helpers.PrintfErrorExit("Error occurred, are you sure your application has the 'Advanced tutor' role?\n")
			default:
				// This will print the status code and body
				helpers.GetRespBody(resp)
		}
	},
}

func init() {
	usersCmd.AddCommand(createCmd)

	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
