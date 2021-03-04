package platformclientv2
import (
	"encoding/json"
)

// Patchexternalsegment
type Patchexternalsegment struct { 
	// Name - Name for the external segment in the system where it originates from.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchexternalsegment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
