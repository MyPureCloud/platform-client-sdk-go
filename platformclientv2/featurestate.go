package platformclientv2
import (
	"encoding/json"
)

// Featurestate
type Featurestate struct { 
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

}

// String returns a JSON representation of the model
func (o *Featurestate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
