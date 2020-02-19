package platformclientv2
import (
	"encoding/json"
)

// Memberentity
type Memberentity struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Memberentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
