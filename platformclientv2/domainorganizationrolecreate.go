package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainorganizationrolecreate
type Domainorganizationrolecreate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The role name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// DefaultRoleId
	DefaultRoleId *string `json:"defaultRoleId,omitempty"`


	// Permissions
	Permissions *[]string `json:"permissions,omitempty"`


	// UnusedPermissions - A collection of the permissions the role is not using
	UnusedPermissions *[]string `json:"unusedPermissions,omitempty"`


	// PermissionPolicies
	PermissionPolicies *[]Domainpermissionpolicy `json:"permissionPolicies,omitempty"`


	// UserCount
	UserCount *int `json:"userCount,omitempty"`


	// RoleNeedsUpdate - Optional unless patch operation.
	RoleNeedsUpdate *bool `json:"roleNeedsUpdate,omitempty"`


	// VarDefault
	VarDefault *bool `json:"default,omitempty"`


	// Base
	Base *bool `json:"base,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Domainorganizationrolecreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainorganizationrolecreate

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultRoleId *string `json:"defaultRoleId,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		UnusedPermissions *[]string `json:"unusedPermissions,omitempty"`
		
		PermissionPolicies *[]Domainpermissionpolicy `json:"permissionPolicies,omitempty"`
		
		UserCount *int `json:"userCount,omitempty"`
		
		RoleNeedsUpdate *bool `json:"roleNeedsUpdate,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		
		Base *bool `json:"base,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		DefaultRoleId: u.DefaultRoleId,
		
		Permissions: u.Permissions,
		
		UnusedPermissions: u.UnusedPermissions,
		
		PermissionPolicies: u.PermissionPolicies,
		
		UserCount: u.UserCount,
		
		RoleNeedsUpdate: u.RoleNeedsUpdate,
		
		VarDefault: u.VarDefault,
		
		Base: u.Base,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainorganizationrolecreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
