package platformclientv2
import (
	"encoding/json"
)

// Servicegoaltemplatelist - List of service goal templates
type Servicegoaltemplatelist struct { 
	// Entities
	Entities *[]Servicegoaltemplate `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatelist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
