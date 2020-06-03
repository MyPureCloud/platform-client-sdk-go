package platformclientv2
import (
	"encoding/json"
)

// Detectednamedentity
type Detectednamedentity struct { 
	// Name - The name of the detected named entity.
	Name *string `json:"name,omitempty"`


	// EntityType - The type of the detected named entity.
	EntityType *string `json:"entityType,omitempty"`


	// Probability - The probability of the detected named entity.
	Probability *float64 `json:"probability,omitempty"`


	// Value - The value of the detected named entity.
	Value *Detectednamedentityvalue `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Detectednamedentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
