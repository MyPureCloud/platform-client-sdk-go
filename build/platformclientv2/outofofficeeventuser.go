package platformclientv2
import (
	"encoding/json"
)

// Outofofficeeventuser
type Outofofficeeventuser struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outofofficeeventuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
