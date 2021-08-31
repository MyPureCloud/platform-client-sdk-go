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

func (o *Tokeninfo) MarshalJSON() ([]byte, error) {
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
		Organization: o.Organization,
		
		HomeOrganization: o.HomeOrganization,
		
		AuthorizedScope: o.AuthorizedScope,
		
		ClonedUser: o.ClonedUser,
		
		OAuthClient: o.OAuthClient,
		Alias:    (*Alias)(o),
	})
}

func (o *Tokeninfo) UnmarshalJSON(b []byte) error {
	var TokeninfoMap map[string]interface{}
	err := json.Unmarshal(b, &TokeninfoMap)
	if err != nil {
		return err
	}
	
	if Organization, ok := TokeninfoMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if HomeOrganization, ok := TokeninfoMap["homeOrganization"].(map[string]interface{}); ok {
		HomeOrganizationString, _ := json.Marshal(HomeOrganization)
		json.Unmarshal(HomeOrganizationString, &o.HomeOrganization)
	}
	
	if AuthorizedScope, ok := TokeninfoMap["authorizedScope"].([]interface{}); ok {
		AuthorizedScopeString, _ := json.Marshal(AuthorizedScope)
		json.Unmarshal(AuthorizedScopeString, &o.AuthorizedScope)
	}
	
	if ClonedUser, ok := TokeninfoMap["clonedUser"].(map[string]interface{}); ok {
		ClonedUserString, _ := json.Marshal(ClonedUser)
		json.Unmarshal(ClonedUserString, &o.ClonedUser)
	}
	
	if OAuthClient, ok := TokeninfoMap["OAuthClient"].(map[string]interface{}); ok {
		OAuthClientString, _ := json.Marshal(OAuthClient)
		json.Unmarshal(OAuthClientString, &o.OAuthClient)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Tokeninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
