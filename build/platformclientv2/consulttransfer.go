package platformclientv2
import (
	"encoding/json"
)

// Consulttransfer
type Consulttransfer struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking. Defaults to DESTINATION
	SpeakTo *string `json:"speakTo,omitempty"`


	// Destination - Destination phone number and name.
	Destination *Destination `json:"destination,omitempty"`

}

// String returns a JSON representation of the model
func (o *Consulttransfer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
