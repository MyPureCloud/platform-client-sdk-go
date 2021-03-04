package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Domainorganizationrolecreate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
