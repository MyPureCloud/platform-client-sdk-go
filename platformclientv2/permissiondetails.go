package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Permissiondetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
