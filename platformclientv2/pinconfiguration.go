package platformclientv2
import (
	"encoding/json"
)

// Pinconfiguration
type Pinconfiguration struct { 
	// MinimumLength
	MinimumLength *int32 `json:"minimumLength,omitempty"`


	// MaximumLength
	MaximumLength *int32 `json:"maximumLength,omitempty"`

}

// String returns a JSON representation of the model
func (o *Pinconfiguration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
