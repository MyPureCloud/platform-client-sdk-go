package platformclientv2
import (
	"encoding/json"
)

// Oauthclientrequest
type Oauthclientrequest struct { 
	// Name - The name of the OAuth client.
	Name *string `json:"name,omitempty"`


	// AccessTokenValiditySeconds - The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds *int64 `json:"accessTokenValiditySeconds,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// RegisteredRedirectUri - List of allowed callbacks for this client. For example: https://myap.example.com/auth/callback
	RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`


	// RoleIds - Deprecated. Use roleDivisions instead.
	RoleIds *[]string `json:"roleIds,omitempty"`


	// AuthorizedGrantType - The OAuth Grant/Client type supported by this client. Code Authorization Grant/Client type - Preferred client type where the Client ID and Secret are required to create tokens. Used where the secret can be secured. Implicit grant type - Client ID only is required to create tokens. Used in browser and mobile apps where the secret can not be secured. SAML2-Bearer extension grant type - SAML2 assertion provider for user authentication at the token endpoint. Client Credential grant type - Used to created access tokens that are tied only to the client. 
	AuthorizedGrantType *string `json:"authorizedGrantType,omitempty"`


	// Scope - The scope requested by this client. Scopes only apply to clients not using the client_credential grant
	Scope *[]string `json:"scope,omitempty"`


	// RoleDivisions - Set of roles and their corresponding divisions associated with this client. Roles and divisions only apply to clients using the client_credential grant
	RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`


	// State - The state of the OAuth client. Active: The OAuth client can be used to create access tokens. This is the default state. Disabled: Access tokens created by the client are invalid and new ones cannot be created. Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthclientrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
