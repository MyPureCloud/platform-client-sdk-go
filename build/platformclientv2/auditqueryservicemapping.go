package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryservicemapping
type Auditqueryservicemapping struct { 
	// Services - List of Services
	Services *[]Auditqueryservice `json:"services,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryservicemapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
