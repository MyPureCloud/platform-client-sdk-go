package platformclientv2
import (
	"encoding/json"
)

// Timeoffrequestentitylist
type Timeoffrequestentitylist struct { 
	// Entities
	Entities *[]Timeoffrequestresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestentitylist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
