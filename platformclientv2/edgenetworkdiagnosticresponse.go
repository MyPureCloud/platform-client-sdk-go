package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgenetworkdiagnosticresponse
type Edgenetworkdiagnosticresponse struct { 
	// CommandCorrelationId - UUID of each executed command on edge
	CommandCorrelationId *string `json:"commandCorrelationId,omitempty"`


	// Diagnostics - Response string of executed command from edge
	Diagnostics *string `json:"diagnostics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnosticresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
