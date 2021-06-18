package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Predictorworkloadbalancing
type Predictorworkloadbalancing struct { 
	// Enabled - Flag to activate and deactivate workload balancing.
	Enabled *bool `json:"enabled,omitempty"`

}

// String returns a JSON representation of the model
func (o *Predictorworkloadbalancing) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
