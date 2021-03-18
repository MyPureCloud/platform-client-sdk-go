package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Mediasummary
type Mediasummary struct { 
	// ContactCenter
	ContactCenter *Mediasummarydetail `json:"contactCenter,omitempty"`


	// Enterprise
	Enterprise *Mediasummarydetail `json:"enterprise,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediasummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
