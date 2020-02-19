package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Domainorgroledifference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
