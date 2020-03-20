package platformclientv2
import (
	"encoding/json"
)

// Digitlength
type Digitlength struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`

}

// String returns a JSON representation of the model
func (o *Digitlength) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
