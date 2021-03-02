package platformclientv2
import (
	"encoding/json"
)

// Contextintent
type Contextintent struct { 
	// Name - The name of the intent.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contextintent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
