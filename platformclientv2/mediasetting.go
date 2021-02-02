package platformclientv2
import (
	"encoding/json"
)

// Mediasetting
type Mediasetting struct { 
	// AlertingTimeoutSeconds
	AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`


	// ServiceLevel
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediasetting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
