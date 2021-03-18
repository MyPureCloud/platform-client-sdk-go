package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queuemessagingaddresses
type Queuemessagingaddresses struct { 
	// SmsAddress
	SmsAddress *Domainentityref `json:"smsAddress,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queuemessagingaddresses) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
