package platformclientv2
import (
	"encoding/json"
)

// Integrationconfigurationinfo - Configuration information for the integration
type Integrationconfigurationinfo struct { 
	// Current - The current, active configuration for the integration.
	Current *Integrationconfiguration `json:"current,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integrationconfigurationinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
