package platformclientv2
import (
	"encoding/json"
)

// Workspacecreate
type Workspacecreate struct { 
	// Name - The workspace name
	Name *string `json:"name,omitempty"`


	// Bucket
	Bucket *string `json:"bucket,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workspacecreate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
