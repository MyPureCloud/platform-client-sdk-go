package platformclientv2
import (
	"encoding/json"
)

// Architectpromptnotificationerrormessageparams
type Architectpromptnotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
