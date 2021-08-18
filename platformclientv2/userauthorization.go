package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Userauthorization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userauthorization

	

	return json.Marshal(&struct { 
		Roles *[]Domainrole `json:"roles,omitempty"`
		
		UnusedRoles *[]Domainrole `json:"unusedRoles,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		PermissionPolicies *[]Resourcepermissionpolicy `json:"permissionPolicies,omitempty"`
		*Alias
	}{ 
		Roles: u.Roles,
		
		UnusedRoles: u.UnusedRoles,
		
		Permissions: u.Permissions,
		
		PermissionPolicies: u.PermissionPolicies,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
