package platformclientv2
import (
	"encoding/json"
)

// Wfmuserentitylisting
type Wfmuserentitylisting struct { 
	// Entities
	Entities *[]User `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
