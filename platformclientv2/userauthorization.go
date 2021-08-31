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

func (o *Userauthorization) MarshalJSON() ([]byte, error) {
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
		Roles: o.Roles,
		
		UnusedRoles: o.UnusedRoles,
		
		Permissions: o.Permissions,
		
		PermissionPolicies: o.PermissionPolicies,
		Alias:    (*Alias)(o),
	})
}

func (o *Userauthorization) UnmarshalJSON(b []byte) error {
	var UserauthorizationMap map[string]interface{}
	err := json.Unmarshal(b, &UserauthorizationMap)
	if err != nil {
		return err
	}
	
	if Roles, ok := UserauthorizationMap["roles"].([]interface{}); ok {
		RolesString, _ := json.Marshal(Roles)
		json.Unmarshal(RolesString, &o.Roles)
	}
	
	if UnusedRoles, ok := UserauthorizationMap["unusedRoles"].([]interface{}); ok {
		UnusedRolesString, _ := json.Marshal(UnusedRoles)
		json.Unmarshal(UnusedRolesString, &o.UnusedRoles)
	}
	
	if Permissions, ok := UserauthorizationMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if PermissionPolicies, ok := UserauthorizationMap["permissionPolicies"].([]interface{}); ok {
		PermissionPoliciesString, _ := json.Marshal(PermissionPolicies)
		json.Unmarshal(PermissionPoliciesString, &o.PermissionPolicies)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
