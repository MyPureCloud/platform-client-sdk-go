package platformclientv2
import (
	"encoding/json"
)

// Updateactioninput
type Updateactioninput struct { 
	// Category - Category of action
	Category *string `json:"category,omitempty"`


	// Name - Name of action
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Version - Version of this action
	Version *int32 `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateactioninput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
