package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref - A reference for an SmsPhoneNumber
type Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref struct { 
	// PhoneNumber - The unique phone number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref
	
	return json.Marshal(&struct { 
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		*Alias
	}{ 
		PhoneNumber: o.PhoneNumber,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangesmsphonenumberrefMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangesmsphonenumberrefMap)
	if err != nil {
		return err
	}
	
	if PhoneNumber, ok := OutboundmessagingmessagingcampaignconfigchangesmsphonenumberrefMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
