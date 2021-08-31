package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthclientrequest
type Oauthclientrequest struct { 
	// Name - The name of the OAuth client.
	Name *string `json:"name,omitempty"`


	// AccessTokenValiditySeconds - The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// RegisteredRedirectUri - List of allowed callbacks for this client. For example: https://myap.example.com/auth/callback
	RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`


	// RoleIds - Deprecated. Use roleDivisions instead.
	RoleIds *[]string `json:"roleIds,omitempty"`


	// AuthorizedGrantType - The OAuth Grant/Client type supported by this client. Code Authorization Grant/Client type - Preferred client type where the Client ID and Secret are required to create tokens. Used where the secret can be secured. PKCE-Enabled Code Authorization grant type - Code grant type which requires PKCE challenge and verifier to create tokens. Used in public clients for increased security. Implicit grant type - Client ID only is required to create tokens. Used in browser and mobile apps where the secret can not be secured. SAML2-Bearer extension grant type - SAML2 assertion provider for user authentication at the token endpoint. Client Credential grant type - Used to created access tokens that are tied only to the client. 
	AuthorizedGrantType *string `json:"authorizedGrantType,omitempty"`


	// Scope - The scope requested by this client. Scopes only apply to clients not using the client_credential grant
	Scope *[]string `json:"scope,omitempty"`


	// RoleDivisions - Set of roles and their corresponding divisions associated with this client. Roles and divisions only apply to clients using the client_credential grant
	RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`


	// State - The state of the OAuth client. Active: The OAuth client can be used to create access tokens. This is the default state. Disabled: Access tokens created by the client are invalid and new ones cannot be created. Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
	State *string `json:"state,omitempty"`


	// DateToDelete - The time at which this client will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateToDelete *time.Time `json:"dateToDelete,omitempty"`

}

func (o *Oauthclientrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthclientrequest
	
	DateToDelete := new(string)
	if o.DateToDelete != nil {
		
		*DateToDelete = timeutil.Strftime(o.DateToDelete, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateToDelete = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`
		
		RoleIds *[]string `json:"roleIds,omitempty"`
		
		AuthorizedGrantType *string `json:"authorizedGrantType,omitempty"`
		
		Scope *[]string `json:"scope,omitempty"`
		
		RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateToDelete *string `json:"dateToDelete,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		AccessTokenValiditySeconds: o.AccessTokenValiditySeconds,
		
		Description: o.Description,
		
		RegisteredRedirectUri: o.RegisteredRedirectUri,
		
		RoleIds: o.RoleIds,
		
		AuthorizedGrantType: o.AuthorizedGrantType,
		
		Scope: o.Scope,
		
		RoleDivisions: o.RoleDivisions,
		
		State: o.State,
		
		DateToDelete: DateToDelete,
		Alias:    (*Alias)(o),
	})
}

func (o *Oauthclientrequest) UnmarshalJSON(b []byte) error {
	var OauthclientrequestMap map[string]interface{}
	err := json.Unmarshal(b, &OauthclientrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := OauthclientrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if AccessTokenValiditySeconds, ok := OauthclientrequestMap["accessTokenValiditySeconds"].(float64); ok {
		AccessTokenValiditySecondsInt := int(AccessTokenValiditySeconds)
		o.AccessTokenValiditySeconds = &AccessTokenValiditySecondsInt
	}
	
	if Description, ok := OauthclientrequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if RegisteredRedirectUri, ok := OauthclientrequestMap["registeredRedirectUri"].([]interface{}); ok {
		RegisteredRedirectUriString, _ := json.Marshal(RegisteredRedirectUri)
		json.Unmarshal(RegisteredRedirectUriString, &o.RegisteredRedirectUri)
	}
	
	if RoleIds, ok := OauthclientrequestMap["roleIds"].([]interface{}); ok {
		RoleIdsString, _ := json.Marshal(RoleIds)
		json.Unmarshal(RoleIdsString, &o.RoleIds)
	}
	
	if AuthorizedGrantType, ok := OauthclientrequestMap["authorizedGrantType"].(string); ok {
		o.AuthorizedGrantType = &AuthorizedGrantType
	}
	
	if Scope, ok := OauthclientrequestMap["scope"].([]interface{}); ok {
		ScopeString, _ := json.Marshal(Scope)
		json.Unmarshal(ScopeString, &o.Scope)
	}
	
	if RoleDivisions, ok := OauthclientrequestMap["roleDivisions"].([]interface{}); ok {
		RoleDivisionsString, _ := json.Marshal(RoleDivisions)
		json.Unmarshal(RoleDivisionsString, &o.RoleDivisions)
	}
	
	if State, ok := OauthclientrequestMap["state"].(string); ok {
		o.State = &State
	}
	
	if dateToDeleteString, ok := OauthclientrequestMap["dateToDelete"].(string); ok {
		DateToDelete, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateToDeleteString)
		o.DateToDelete = &DateToDelete
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Oauthclientrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
