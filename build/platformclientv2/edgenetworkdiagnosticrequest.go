package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgenetworkdiagnosticrequest
type Edgenetworkdiagnosticrequest struct { 
	// Host - IPv4/6 address or host to be probed for connectivity. No port allowed.
	Host *string `json:"host,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnosticrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
