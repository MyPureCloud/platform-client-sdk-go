package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthclientlisting
type Oauthclientlisting struct { 
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

func (u *Oauthclientlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthclientlisting

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateToDelete := new(string)
	if u.DateToDelete != nil {
		
		*DateToDelete = timeutil.Strftime(u.DateToDelete, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		Scope *[]string `json:"scope,omitempty"`
		
		RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateToDelete *string `json:"dateToDelete,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		AccessTokenValiditySeconds: u.AccessTokenValiditySeconds,
		
		Description: u.Description,
		
		RegisteredRedirectUri: u.RegisteredRedirectUri,
		
		Secret: u.Secret,
		
		RoleIds: u.RoleIds,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: u.CreatedBy,
		
		ModifiedBy: u.ModifiedBy,
		
		Scope: u.Scope,
		
		RoleDivisions: u.RoleDivisions,
		
		State: u.State,
		
		DateToDelete: DateToDelete,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Oauthclientlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
