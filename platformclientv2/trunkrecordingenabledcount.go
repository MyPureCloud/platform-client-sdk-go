package platformclientv2
import (
	"encoding/json"
)

// Trunkrecordingenabledcount
type Trunkrecordingenabledcount struct { 
	// EnabledCount - The amount of trunks that have recording enabled
	EnabledCount *int `json:"enabledCount,omitempty"`


	// DisabledCount - The amount of trunks that do not have recording enabled
	DisabledCount *int `json:"disabledCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkrecordingenabledcount) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
