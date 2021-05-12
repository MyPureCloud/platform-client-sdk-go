package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Automatictimezonemappingsettings
type Automatictimezonemappingsettings struct { 
	// CallableWindows - The time intervals to use for automatic time zone mapping.
	CallableWindows *[]Callablewindow `json:"callableWindows,omitempty"`

}

// String returns a JSON representation of the model
func (o *Automatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
