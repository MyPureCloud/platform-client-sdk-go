package platformclientv2
import (
	"encoding/json"
)

// Userappconfigurationinfo - Configuration information for the integration
type Userappconfigurationinfo struct { 
	// Current - The current, active configuration for the integration.
	Current *Integrationconfiguration `json:"current,omitempty"`


	// Effective - The effective configuration for the app, containing the integration specific configuration along with overrides specified in the integration type.
	Effective *Effectiveconfiguration `json:"effective,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userappconfigurationinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
