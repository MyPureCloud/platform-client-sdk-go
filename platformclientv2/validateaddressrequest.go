package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Validateaddressrequest
type Validateaddressrequest struct { 
	// Address - Address schema
	Address *Streetaddress `json:"address,omitempty"`

}

// String returns a JSON representation of the model
func (o *Validateaddressrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
