package platformclientv2
import (
	"encoding/json"
)

// Bullseye
type Bullseye struct { 
	// Rings
	Rings *[]Ring `json:"rings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bullseye) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
