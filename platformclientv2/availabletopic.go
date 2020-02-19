package platformclientv2
import (
	"encoding/json"
)

// Availabletopic
type Availabletopic struct { 
	// Description
	Description *string `json:"description,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// RequiresPermissions
	RequiresPermissions *[]string `json:"requiresPermissions,omitempty"`


	// Schema
	Schema *map[string]interface{} `json:"schema,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availabletopic) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
