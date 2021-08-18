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

func (u *Permissiondetails) MarshalJSON() ([]byte, error) {
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
		VarType: u.VarType,
		
		Permissions: u.Permissions,
		
		AllowsCurrentUser: u.AllowsCurrentUser,
		
		Enforced: u.Enforced,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Permissiondetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
