package platformclientv2
import (
	"encoding/json"
)

// Consulttransferupdate
type Consulttransferupdate struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking.
	SpeakTo *string `json:"speakTo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Consulttransferupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
