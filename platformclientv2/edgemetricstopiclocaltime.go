package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopiclocaltime
type Edgemetricstopiclocaltime struct { 
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
func (o *Edgemetricstopiclocaltime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
