package platformclientv2
import (
	"encoding/json"
)

// Userauthorization
type Userauthorization struct { 
	// Roles
	Roles *[]Domainrole `json:"roles,omitempty"`


	// UnusedRoles - A collection of the roles the user is not using
	UnusedRoles *[]Domainrole `json:"unusedRoles,omitempty"`


	// Permissions - A collection of the permissions granted by all assigned roles
	Permissions *[]string `json:"permissions,omitempty"`


	// PermissionPolicies - The policies configured for assigned permissions.
	PermissionPolicies *[]Resourcepermissionpolicy `json:"permissionPolicies,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userauthorization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
