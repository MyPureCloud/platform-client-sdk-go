package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainorgroledifference
type Domainorgroledifference struct { 
	// RemovedPermissionPolicies
	RemovedPermissionPolicies *[]Domainpermissionpolicy `json:"removedPermissionPolicies,omitempty"`


	// AddedPermissionPolicies
	AddedPermissionPolicies *[]Domainpermissionpolicy `json:"addedPermissionPolicies,omitempty"`


	// SamePermissionPolicies
	SamePermissionPolicies *[]Domainpermissionpolicy `json:"samePermissionPolicies,omitempty"`


	// UserOrgRole
	UserOrgRole *Domainorganizationrole `json:"userOrgRole,omitempty"`


	// RoleFromDefault
	RoleFromDefault *Domainorganizationrole `json:"roleFromDefault,omitempty"`

}

func (o *Domainorgroledifference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainorgroledifference
	
	return json.Marshal(&struct { 
		RemovedPermissionPolicies *[]Domainpermissionpolicy `json:"removedPermissionPolicies,omitempty"`
		
		AddedPermissionPolicies *[]Domainpermissionpolicy `json:"addedPermissionPolicies,omitempty"`
		
		SamePermissionPolicies *[]Domainpermissionpolicy `json:"samePermissionPolicies,omitempty"`
		
		UserOrgRole *Domainorganizationrole `json:"userOrgRole,omitempty"`
		
		RoleFromDefault *Domainorganizationrole `json:"roleFromDefault,omitempty"`
		*Alias
	}{ 
		RemovedPermissionPolicies: o.RemovedPermissionPolicies,
		
		AddedPermissionPolicies: o.AddedPermissionPolicies,
		
		SamePermissionPolicies: o.SamePermissionPolicies,
		
		UserOrgRole: o.UserOrgRole,
		
		RoleFromDefault: o.RoleFromDefault,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainorgroledifference) UnmarshalJSON(b []byte) error {
	var DomainorgroledifferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DomainorgroledifferenceMap)
	if err != nil {
		return err
	}
	
	if RemovedPermissionPolicies, ok := DomainorgroledifferenceMap["removedPermissionPolicies"].([]interface{}); ok {
		RemovedPermissionPoliciesString, _ := json.Marshal(RemovedPermissionPolicies)
		json.Unmarshal(RemovedPermissionPoliciesString, &o.RemovedPermissionPolicies)
	}
	
	if AddedPermissionPolicies, ok := DomainorgroledifferenceMap["addedPermissionPolicies"].([]interface{}); ok {
		AddedPermissionPoliciesString, _ := json.Marshal(AddedPermissionPolicies)
		json.Unmarshal(AddedPermissionPoliciesString, &o.AddedPermissionPolicies)
	}
	
	if SamePermissionPolicies, ok := DomainorgroledifferenceMap["samePermissionPolicies"].([]interface{}); ok {
		SamePermissionPoliciesString, _ := json.Marshal(SamePermissionPolicies)
		json.Unmarshal(SamePermissionPoliciesString, &o.SamePermissionPolicies)
	}
	
	if UserOrgRole, ok := DomainorgroledifferenceMap["userOrgRole"].(map[string]interface{}); ok {
		UserOrgRoleString, _ := json.Marshal(UserOrgRole)
		json.Unmarshal(UserOrgRoleString, &o.UserOrgRole)
	}
	
	if RoleFromDefault, ok := DomainorgroledifferenceMap["roleFromDefault"].(map[string]interface{}); ok {
		RoleFromDefaultString, _ := json.Marshal(RoleFromDefault)
		json.Unmarshal(RoleFromDefaultString, &o.RoleFromDefault)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainorgroledifference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
