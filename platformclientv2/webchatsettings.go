package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatsettings
type Webchatsettings struct { 
	// RequireDeployment
	RequireDeployment *bool `json:"requireDeployment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
