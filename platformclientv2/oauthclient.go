package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthclient
type Oauthclient struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the OAuth client.
	Name *string `json:"name,omitempty"`


	// AccessTokenValiditySeconds - The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// RegisteredRedirectUri - List of allowed callbacks for this client. For example: https://myap.example.com/auth/callback
	RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`


	// Secret - System created secret assigned to this client. Secrets are required for code authorization and client credential grants.
	Secret *string `json:"secret,omitempty"`


	// RoleIds - Deprecated. Use roleDivisions instead.
	RoleIds *[]string `json:"roleIds,omitempty"`


	// DateCreated - Date this client was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this client was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User that created this client
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User that last modified this client
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Oauthclient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthclient
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateToDelete := new(string)
	if o.DateToDelete != nil {
		
		*DateToDelete = timeutil.Strftime(o.DateToDelete, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateToDelete = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`
		
		Secret *string `json:"secret,omitempty"`
		
		RoleIds *[]string `json:"roleIds,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		AuthorizedGrantType *string `json:"authorizedGrantType,omitempty"`
		
		Scope *[]string `json:"scope,omitempty"`
		
		RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateToDelete *string `json:"dateToDelete,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		AccessTokenValiditySeconds: o.AccessTokenValiditySeconds,
		
		Description: o.Description,
		
		RegisteredRedirectUri: o.RegisteredRedirectUri,
		
		Secret: o.Secret,
		
		RoleIds: o.RoleIds,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		AuthorizedGrantType: o.AuthorizedGrantType,
		
		Scope: o.Scope,
		
		RoleDivisions: o.RoleDivisions,
		
		State: o.State,
		
		DateToDelete: DateToDelete,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Oauthclient) UnmarshalJSON(b []byte) error {
	var OauthclientMap map[string]interface{}
	err := json.Unmarshal(b, &OauthclientMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OauthclientMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OauthclientMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if AccessTokenValiditySeconds, ok := OauthclientMap["accessTokenValiditySeconds"].(float64); ok {
		AccessTokenValiditySecondsInt := int(AccessTokenValiditySeconds)
		o.AccessTokenValiditySeconds = &AccessTokenValiditySecondsInt
	}
	
	if Description, ok := OauthclientMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if RegisteredRedirectUri, ok := OauthclientMap["registeredRedirectUri"].([]interface{}); ok {
		RegisteredRedirectUriString, _ := json.Marshal(RegisteredRedirectUri)
		json.Unmarshal(RegisteredRedirectUriString, &o.RegisteredRedirectUri)
	}
	
	if Secret, ok := OauthclientMap["secret"].(string); ok {
		o.Secret = &Secret
	}
    
	if RoleIds, ok := OauthclientMap["roleIds"].([]interface{}); ok {
		RoleIdsString, _ := json.Marshal(RoleIds)
		json.Unmarshal(RoleIdsString, &o.RoleIds)
	}
	
	if dateCreatedString, ok := OauthclientMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OauthclientMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := OauthclientMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := OauthclientMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if AuthorizedGrantType, ok := OauthclientMap["authorizedGrantType"].(string); ok {
		o.AuthorizedGrantType = &AuthorizedGrantType
	}
    
	if Scope, ok := OauthclientMap["scope"].([]interface{}); ok {
		ScopeString, _ := json.Marshal(Scope)
		json.Unmarshal(ScopeString, &o.Scope)
	}
	
	if RoleDivisions, ok := OauthclientMap["roleDivisions"].([]interface{}); ok {
		RoleDivisionsString, _ := json.Marshal(RoleDivisions)
		json.Unmarshal(RoleDivisionsString, &o.RoleDivisions)
	}
	
	if State, ok := OauthclientMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateToDeleteString, ok := OauthclientMap["dateToDelete"].(string); ok {
		DateToDelete, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateToDeleteString)
		o.DateToDelete = &DateToDelete
	}
	
	if SelfUri, ok := OauthclientMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Oauthclient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
