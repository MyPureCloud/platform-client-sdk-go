package platformclientv2
import (
	"encoding/json"
)

// Architectflownotificationerrormessageparams
type Architectflownotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflownotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
