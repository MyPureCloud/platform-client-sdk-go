package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainorganizationroleupdate
type Domainorganizationroleupdate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the role
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


	// Base
	Base *bool `json:"base,omitempty"`


	// VarDefault
	VarDefault *bool `json:"default,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Domainorganizationroleupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainorganizationroleupdate
	
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
		
		Base *bool `json:"base,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		
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
		
		Base: o.Base,
		
		VarDefault: o.VarDefault,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainorganizationroleupdate) UnmarshalJSON(b []byte) error {
	var DomainorganizationroleupdateMap map[string]interface{}
	err := json.Unmarshal(b, &DomainorganizationroleupdateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainorganizationroleupdateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DomainorganizationroleupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := DomainorganizationroleupdateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DefaultRoleId, ok := DomainorganizationroleupdateMap["defaultRoleId"].(string); ok {
		o.DefaultRoleId = &DefaultRoleId
	}
    
	if Permissions, ok := DomainorganizationroleupdateMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if UnusedPermissions, ok := DomainorganizationroleupdateMap["unusedPermissions"].([]interface{}); ok {
		UnusedPermissionsString, _ := json.Marshal(UnusedPermissions)
		json.Unmarshal(UnusedPermissionsString, &o.UnusedPermissions)
	}
	
	if PermissionPolicies, ok := DomainorganizationroleupdateMap["permissionPolicies"].([]interface{}); ok {
		PermissionPoliciesString, _ := json.Marshal(PermissionPolicies)
		json.Unmarshal(PermissionPoliciesString, &o.PermissionPolicies)
	}
	
	if UserCount, ok := DomainorganizationroleupdateMap["userCount"].(float64); ok {
		UserCountInt := int(UserCount)
		o.UserCount = &UserCountInt
	}
	
	if RoleNeedsUpdate, ok := DomainorganizationroleupdateMap["roleNeedsUpdate"].(bool); ok {
		o.RoleNeedsUpdate = &RoleNeedsUpdate
	}
    
	if Base, ok := DomainorganizationroleupdateMap["base"].(bool); ok {
		o.Base = &Base
	}
    
	if VarDefault, ok := DomainorganizationroleupdateMap["default"].(bool); ok {
		o.VarDefault = &VarDefault
	}
    
	if SelfUri, ok := DomainorganizationroleupdateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainorganizationroleupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
