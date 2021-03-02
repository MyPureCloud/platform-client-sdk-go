package platformclientv2
import (
	"encoding/json"
)

// Digits
type Digits struct { 
	// Digits - A string representing the digits pressed on phone.
	Digits *string `json:"digits,omitempty"`

}

// String returns a JSON representation of the model
func (o *Digits) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
