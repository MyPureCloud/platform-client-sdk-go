package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref
type Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref struct { 
	// PhoneNumber
	PhoneNumber *string `json:"phoneNumber,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
