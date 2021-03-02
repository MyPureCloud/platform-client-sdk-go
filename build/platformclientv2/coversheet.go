package platformclientv2
import (
	"encoding/json"
)

// Coversheet
type Coversheet struct { 
	// Notes - Text to be added to the coversheet
	Notes *string `json:"notes,omitempty"`


	// Locale - Locale, e.g. = en-US
	Locale *string `json:"locale,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coversheet) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
