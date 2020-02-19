package platformclientv2
import (
	"encoding/json"
)

// Usergreetingeventgreetingowner
type Usergreetingeventgreetingowner struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usergreetingeventgreetingowner) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
