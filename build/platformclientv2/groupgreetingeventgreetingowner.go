package platformclientv2
import (
	"encoding/json"
)

// Groupgreetingeventgreetingowner
type Groupgreetingeventgreetingowner struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingowner) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
