package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustmembercreate
type Trustmembercreate struct { 
	// Id - Trustee User or Group Id
	Id *string `json:"id,omitempty"`


	// RoleIds - The list of roles to be granted to this user or group. Roles will be granted in all divisions.
	RoleIds *[]string `json:"roleIds,omitempty"`


	// RoleDivisions - The list of trustor organization roles granting this user or group access paired with the divisions for those roles.
	RoleDivisions *Roledivisiongrants `json:"roleDivisions,omitempty"`

}

func (u *Trustmembercreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustmembercreate

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		RoleIds *[]string `json:"roleIds,omitempty"`
		
		RoleDivisions *Roledivisiongrants `json:"roleDivisions,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		RoleIds: u.RoleIds,
		
		RoleDivisions: u.RoleDivisions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trustmembercreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
