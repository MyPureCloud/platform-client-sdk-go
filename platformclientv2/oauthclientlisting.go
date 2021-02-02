package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Oauthclientlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
