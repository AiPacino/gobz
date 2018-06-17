package ginbz

import (
	"encoding/json"
	"net/http"

	"github.com/bigzhu/gobz/httpbz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleUserInfo 用户信息
type GoogleUserInfo struct {
	Email string `json:"email"`
}

// OauthGoogle oauth2
func OauthGoogle(outc interface{}, redirectURL string, clientID string, clientSecret string) (googleUserInfo GoogleUserInfo, err error) {
	c := outc.(*gin.Context)
	var googleConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	state := c.Query("state")
	if state == "" {
		url := googleConf.AuthCodeURL("state")
		c.Redirect(http.StatusFound, url)
	}
	code := c.Query("code")
	token, err := googleConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ERROR: err.Error()})
		return
	}

	data, err := httpbz.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &googleUserInfo)
	return
}
