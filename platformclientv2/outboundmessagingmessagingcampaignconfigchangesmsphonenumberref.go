package platformclientv2
import (
	"encoding/json"
)

// Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref
type Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref struct { 
	// PhoneNumber
	PhoneNumber *string `json:"phoneNumber,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
