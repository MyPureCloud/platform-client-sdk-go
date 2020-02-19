package platformclientv2
import (
	"encoding/json"
)

// Facebookintegrationrequest
type Facebookintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Facebook Integration
	Name *string `json:"name,omitempty"`


	// PageAccessToken - The long-lived Page Access Token of a facebook page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When a pageAccessToken is provided, pageId and userAccessToken are not required.
	PageAccessToken *string `json:"pageAccessToken,omitempty"`


	// UserAccessToken - The short-lived User Access Token of the facebook user logged into the facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When userAccessToken is provided, pageId is mandatory.  When userAccessToken/pageId combination is provided, pageAccessToken is not required.
	UserAccessToken *string `json:"userAccessToken,omitempty"`


	// PageId - The page Id of a facebook page. The pageId is required when userAccessToken is provided.
	PageId *string `json:"pageId,omitempty"`


	// AppId - The app Id of a facebook app
	AppId *string `json:"appId,omitempty"`


	// AppSecret - The app Secret of a facebook app
	AppSecret *string `json:"appSecret,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facebookintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
