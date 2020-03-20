package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradaymetric
type Wfmintradaydataupdatetopicintradaymetric struct { 
	// Category
	Category *string `json:"category,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaymetric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
