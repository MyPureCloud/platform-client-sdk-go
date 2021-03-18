package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { 
	// DataAvailabilityDate
	DataAvailabilityDate *Conversationdetailsdatalakeavailabilitytopicdatetime `json:"dataAvailabilityDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
