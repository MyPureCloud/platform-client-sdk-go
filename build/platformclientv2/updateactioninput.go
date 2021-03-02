package platformclientv2
import (
	"encoding/json"
)

// Updateactioninput
type Updateactioninput struct { 
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`


	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Version - Version of this action
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateactioninput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
