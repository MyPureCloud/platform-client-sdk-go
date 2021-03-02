package platformclientv2
import (
	"encoding/json"
)

// Intentfeedback
type Intentfeedback struct { 
	// Name - The name of the detected intent.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected intent.
	Probability *float64 `json:"probability,omitempty"`


	// Entities - The collection of named entities detected.
	Entities *[]Detectednamedentity `json:"entities,omitempty"`


	// Assessment - The assessment on the detection for feedback text.
	Assessment *string `json:"assessment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intentfeedback) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
