package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Permissiondetails
type Permissiondetails struct { 
	// VarType - The type of permission requirement
	VarType *string `json:"type,omitempty"`


	// Permissions - List of required permissions
	Permissions *[]string `json:"permissions,omitempty"`


	// AllowsCurrentUser - Whether the current user can subscribe, when division permissions are otherwise required
	AllowsCurrentUser *bool `json:"allowsCurrentUser,omitempty"`


	// Enforced - Whether or not this permission requirement is enforced
	Enforced *bool `json:"enforced,omitempty"`

}

func (o *Permissiondetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Permissiondetails
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		AllowsCurrentUser *bool `json:"allowsCurrentUser,omitempty"`
		
		Enforced *bool `json:"enforced,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Permissions: o.Permissions,
		
		AllowsCurrentUser: o.AllowsCurrentUser,
		
		Enforced: o.Enforced,
		Alias:    (*Alias)(o),
	})
}

func (o *Permissiondetails) UnmarshalJSON(b []byte) error {
	var PermissiondetailsMap map[string]interface{}
	err := json.Unmarshal(b, &PermissiondetailsMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := PermissiondetailsMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Permissions, ok := PermissiondetailsMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if AllowsCurrentUser, ok := PermissiondetailsMap["allowsCurrentUser"].(bool); ok {
		o.AllowsCurrentUser = &AllowsCurrentUser
	}
    
	if Enforced, ok := PermissiondetailsMap["enforced"].(bool); ok {
		o.Enforced = &Enforced
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Permissiondetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
