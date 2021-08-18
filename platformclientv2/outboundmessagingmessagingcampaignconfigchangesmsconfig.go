package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Outboundmessagingmessagingcampaignconfigchangesmsconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangesmsconfig

	

	return json.Marshal(&struct { 
		MessageColumn *string `json:"messageColumn,omitempty"`
		
		PhoneColumn *string `json:"phoneColumn,omitempty"`
		
		SenderSmsPhoneNumber *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref `json:"senderSmsPhoneNumber,omitempty"`
		
		ContentTemplate *Outboundmessagingmessagingcampaignconfigchangeresponseref `json:"contentTemplate,omitempty"`
		*Alias
	}{ 
		MessageColumn: u.MessageColumn,
		
		PhoneColumn: u.PhoneColumn,
		
		SenderSmsPhoneNumber: u.SenderSmsPhoneNumber,
		
		ContentTemplate: u.ContentTemplate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
