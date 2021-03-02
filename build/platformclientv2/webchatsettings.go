package platformclientv2
import (
	"encoding/json"
)

// Webchatsettings
type Webchatsettings struct { 
	// RequireDeployment
	RequireDeployment *bool `json:"requireDeployment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
