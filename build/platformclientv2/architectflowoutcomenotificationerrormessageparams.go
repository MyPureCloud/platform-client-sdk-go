package platformclientv2
import (
	"encoding/json"
)

// Architectflowoutcomenotificationerrormessageparams
type Architectflowoutcomenotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
