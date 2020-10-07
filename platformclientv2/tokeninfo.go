package platformclientv2
import (
	"encoding/json"
)

// Tokeninfo
type Tokeninfo struct { 
	// Organization - The current organization
	Organization *Namedentity `json:"organization,omitempty"`


	// HomeOrganization - The token's home organization
	HomeOrganization *Namedentity `json:"homeOrganization,omitempty"`


	// AuthorizedScope - The list of scopes authorized for the OAuth client
	AuthorizedScope *[]string `json:"authorizedScope,omitempty"`


	// ClonedUser - Only present when a user is a clone of trustee user in the trustor org.
	ClonedUser *Tokeninfocloneduser `json:"clonedUser,omitempty"`


	// OAuthClient
	OAuthClient *Orgoauthclient `json:"OAuthClient,omitempty"`

}

// String returns a JSON representation of the model
func (o *Tokeninfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
