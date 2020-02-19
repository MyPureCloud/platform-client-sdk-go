package platformclientv2
import (
	"encoding/json"
)

// Callforwardingeventuser
type Callforwardingeventuser struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callforwardingeventuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
