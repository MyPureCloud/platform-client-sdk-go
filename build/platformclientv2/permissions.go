package platformclientv2
import (
	"encoding/json"
)

// Permissions
type Permissions struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Ids - List of permission ids.
	Ids *[]string `json:"ids,omitempty"`

}

// String returns a JSON representation of the model
func (o *Permissions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
