package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Tokeninfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Tokeninfo

	

	return json.Marshal(&struct { 
		Organization *Namedentity `json:"organization,omitempty"`
		
		HomeOrganization *Namedentity `json:"homeOrganization,omitempty"`
		
		AuthorizedScope *[]string `json:"authorizedScope,omitempty"`
		
		ClonedUser *Tokeninfocloneduser `json:"clonedUser,omitempty"`
		
		OAuthClient *Orgoauthclient `json:"OAuthClient,omitempty"`
		*Alias
	}{ 
		Organization: u.Organization,
		
		HomeOrganization: u.HomeOrganization,
		
		AuthorizedScope: u.AuthorizedScope,
		
		ClonedUser: u.ClonedUser,
		
		OAuthClient: u.OAuthClient,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Tokeninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
