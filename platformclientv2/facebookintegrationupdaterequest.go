package platformclientv2
import (
	"encoding/json"
)

// Facebookintegrationupdaterequest
type Facebookintegrationupdaterequest struct { 
	// PageAccessToken - The long-lived Page Access Token of a facebook page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
	PageAccessToken *string `json:"pageAccessToken,omitempty"`


	// UserAccessToken - The short-lived User Access Token of the facebook user logged into the facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
	UserAccessToken *string `json:"userAccessToken,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
