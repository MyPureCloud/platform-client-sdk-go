package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Ntpsettings
type Ntpsettings struct { 
	// Servers - List of NTP servers, in priority order
	Servers *[]string `json:"servers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ntpsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
