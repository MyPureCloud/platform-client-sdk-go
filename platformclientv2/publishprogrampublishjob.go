package platformclientv2
import (
	"encoding/json"
)

// Publishprogrampublishjob
type Publishprogrampublishjob struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Publishprogrampublishjob) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
