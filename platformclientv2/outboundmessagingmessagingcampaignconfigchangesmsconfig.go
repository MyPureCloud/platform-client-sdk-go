package platformclientv2
import (
	"encoding/json"
)

// Outboundmessagingmessagingcampaignconfigchangesmsconfig
type Outboundmessagingmessagingcampaignconfigchangesmsconfig struct { 
	// MessageColumn
	MessageColumn *string `json:"messageColumn,omitempty"`


	// PhoneColumn
	PhoneColumn *string `json:"phoneColumn,omitempty"`


	// SenderSmsPhoneNumber
	SenderSmsPhoneNumber *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref `json:"senderSmsPhoneNumber,omitempty"`


	// ContentTemplate
	ContentTemplate *Outboundmessagingmessagingcampaignconfigchangeresponseref `json:"contentTemplate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsconfig) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
