package platformclientv2
import (
	"encoding/json"
)

// Callcommand
type Callcommand struct { 
	// CallNumber - The phone number to dial for this call.
	CallNumber *string `json:"callNumber,omitempty"`


	// PhoneColumn - For a dialer preview or scheduled callback, the phone column associated with the phone number
	PhoneColumn *string `json:"phoneColumn,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callcommand) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
