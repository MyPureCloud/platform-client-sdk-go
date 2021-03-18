package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradelistresponse
type Shifttradelistresponse struct { 
	// Entities
	Entities *[]Shifttraderesponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradelistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
