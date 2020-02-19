package platformclientv2
import (
	"encoding/json"
)

// Trusteeauthorization
type Trusteeauthorization struct { 
	// Permissions - Permissions that the trustee user has in the trustor organization
	Permissions *[]string `json:"permissions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trusteeauthorization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
