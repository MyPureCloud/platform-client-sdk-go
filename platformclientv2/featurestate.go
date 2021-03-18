package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Featurestate
type Featurestate struct { 
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

}

// String returns a JSON representation of the model
func (o *Featurestate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
