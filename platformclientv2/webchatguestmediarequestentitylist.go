package platformclientv2
import (
	"encoding/json"
)

// Webchatguestmediarequestentitylist
type Webchatguestmediarequestentitylist struct { 
	// Entities
	Entities *[]Webchatguestmediarequest `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequestentitylist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
