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

func (o *Trustmembercreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustmembercreate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		RoleIds *[]string `json:"roleIds,omitempty"`
		
		RoleDivisions *Roledivisiongrants `json:"roleDivisions,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		RoleIds: o.RoleIds,
		
		RoleDivisions: o.RoleDivisions,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustmembercreate) UnmarshalJSON(b []byte) error {
	var TrustmembercreateMap map[string]interface{}
	err := json.Unmarshal(b, &TrustmembercreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrustmembercreateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if RoleIds, ok := TrustmembercreateMap["roleIds"].([]interface{}); ok {
		RoleIdsString, _ := json.Marshal(RoleIds)
		json.Unmarshal(RoleIdsString, &o.RoleIds)
	}
	
	if RoleDivisions, ok := TrustmembercreateMap["roleDivisions"].(map[string]interface{}); ok {
		RoleDivisionsString, _ := json.Marshal(RoleDivisions)
		json.Unmarshal(RoleDivisionsString, &o.RoleDivisions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustmembercreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
