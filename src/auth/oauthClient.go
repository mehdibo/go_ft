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
