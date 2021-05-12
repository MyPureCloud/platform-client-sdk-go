package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Locationemergencynumber
type Locationemergencynumber struct { 
	// E164
	E164 *string `json:"e164,omitempty"`


	// Number
	Number *string `json:"number,omitempty"`


	// VarType - The type of emergency number.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Locationemergencynumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
