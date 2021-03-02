package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricstopiclocaltime
type Trunkmetricstopiclocaltime struct { 
	// Hour
	Hour *int `json:"hour,omitempty"`


	// Minute
	Minute *int `json:"minute,omitempty"`


	// Second
	Second *int `json:"second,omitempty"`


	// Nano
	Nano *int `json:"nano,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopiclocaltime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
