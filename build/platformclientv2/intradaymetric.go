package platformclientv2
import (
	"encoding/json"
)

// Intradaymetric
type Intradaymetric struct { 
	// Category - The metric category
	Category *string `json:"category,omitempty"`


	// Version - The current version id for this metric category
	Version *string `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradaymetric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
