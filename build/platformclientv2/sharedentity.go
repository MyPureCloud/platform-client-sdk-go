package platformclientv2
import (
	"encoding/json"
)

// Sharedentity
type Sharedentity struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sharedentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
