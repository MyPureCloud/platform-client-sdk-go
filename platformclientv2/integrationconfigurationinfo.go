package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationconfigurationinfo - Configuration information for the integration
type Integrationconfigurationinfo struct { 
	// Current - The current, active configuration for the integration.
	Current *Integrationconfiguration `json:"current,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integrationconfigurationinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
