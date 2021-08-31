package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainorganizationrole
type Domainorganizationrole struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
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

func (o *Domainorganizationrole) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainorganizationrole
	
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
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DefaultRoleId: o.DefaultRoleId,
		
		Permissions: o.Permissions,
		
		UnusedPermissions: o.UnusedPermissions,
		
		PermissionPolicies: o.PermissionPolicies,
		
		UserCount: o.UserCount,
		
		RoleNeedsUpdate: o.RoleNeedsUpdate,
		
		VarDefault: o.VarDefault,
		
		Base: o.Base,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainorganizationrole) UnmarshalJSON(b []byte) error {
	var DomainorganizationroleMap map[string]interface{}
	err := json.Unmarshal(b, &DomainorganizationroleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainorganizationroleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DomainorganizationroleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := DomainorganizationroleMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if DefaultRoleId, ok := DomainorganizationroleMap["defaultRoleId"].(string); ok {
		o.DefaultRoleId = &DefaultRoleId
	}
	
	if Permissions, ok := DomainorganizationroleMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if UnusedPermissions, ok := DomainorganizationroleMap["unusedPermissions"].([]interface{}); ok {
		UnusedPermissionsString, _ := json.Marshal(UnusedPermissions)
		json.Unmarshal(UnusedPermissionsString, &o.UnusedPermissions)
	}
	
	if PermissionPolicies, ok := DomainorganizationroleMap["permissionPolicies"].([]interface{}); ok {
		PermissionPoliciesString, _ := json.Marshal(PermissionPolicies)
		json.Unmarshal(PermissionPoliciesString, &o.PermissionPolicies)
	}
	
	if UserCount, ok := DomainorganizationroleMap["userCount"].(float64); ok {
		UserCountInt := int(UserCount)
		o.UserCount = &UserCountInt
	}
	
	if RoleNeedsUpdate, ok := DomainorganizationroleMap["roleNeedsUpdate"].(bool); ok {
		o.RoleNeedsUpdate = &RoleNeedsUpdate
	}
	
	if VarDefault, ok := DomainorganizationroleMap["default"].(bool); ok {
		o.VarDefault = &VarDefault
	}
	
	if Base, ok := DomainorganizationroleMap["base"].(bool); ok {
		o.Base = &Base
	}
	
	if SelfUri, ok := DomainorganizationroleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainorganizationrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
