package platformclientv2
import (
	"encoding/json"
)

// Scimv2enterpriseuser - Defines an SCIM enterprise user.
type Scimv2enterpriseuser struct { 
	// Division - The division that the user belongs to.
	Division *string `json:"division,omitempty"`


	// Department - The department that the user belongs to.
	Department *string `json:"department,omitempty"`


	// Manager - The user's manager.
	Manager *Manager `json:"manager,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2enterpriseuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
