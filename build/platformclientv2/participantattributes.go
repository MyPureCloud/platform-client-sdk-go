package platformclientv2
import (
	"encoding/json"
)

// Participantattributes
type Participantattributes struct { 
	// Attributes - The map of attribute keys to values.
	Attributes *map[string]string `json:"attributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Participantattributes) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
