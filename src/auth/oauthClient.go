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
package auth

import (
	"context"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

func GetOauthClient(v *viper.Viper) *http.Client {
	ctx := context.Background()
	oauthConfig := &clientcredentials.Config{
		ClientID:       v.GetString("client_id"),
		ClientSecret:   v.GetString("client_secret"),
		TokenURL:       v.GetString("token_endpoint"),
		Scopes:         v.GetStringSlice("scopes"),
		EndpointParams: nil,
	}
	return oauthConfig.Client(ctx)
}
