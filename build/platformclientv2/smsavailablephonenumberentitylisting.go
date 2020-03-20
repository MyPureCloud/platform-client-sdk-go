package platformclientv2
import (
	"encoding/json"
)

// Smsavailablephonenumberentitylisting
type Smsavailablephonenumberentitylisting struct { 
	// Entities
	Entities *[]Smsavailablephonenumber `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumberentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
