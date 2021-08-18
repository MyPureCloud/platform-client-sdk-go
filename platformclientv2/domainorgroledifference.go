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

func (u *Domainorgroledifference) MarshalJSON() ([]byte, error) {
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
		RemovedPermissionPolicies: u.RemovedPermissionPolicies,
		
		AddedPermissionPolicies: u.AddedPermissionPolicies,
		
		SamePermissionPolicies: u.SamePermissionPolicies,
		
		UserOrgRole: u.UserOrgRole,
		
		RoleFromDefault: u.RoleFromDefault,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainorgroledifference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
