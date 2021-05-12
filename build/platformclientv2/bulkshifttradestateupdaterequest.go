package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkshifttradestateupdaterequest
type Bulkshifttradestateupdaterequest struct { 
	// Entities - The shift trades to update
	Entities *[]Bulkupdateshifttradestaterequestitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkshifttradestateupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
