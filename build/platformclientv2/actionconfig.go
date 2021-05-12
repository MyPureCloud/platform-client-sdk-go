package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Actionconfig - Defines components of the Action Config.
type Actionconfig struct { 
	// Request - Configuration of outbound request.
	Request *Requestconfig `json:"request,omitempty"`


	// Response - Configuration of response processing.
	Response *Responseconfig `json:"response,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actionconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
