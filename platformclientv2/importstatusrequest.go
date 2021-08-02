package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Importstatusrequest
type Importstatusrequest struct { 
	// Status - New status for existing import operation
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Importstatusrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
