package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref
type Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref struct { 
	// PhoneNumber
	PhoneNumber *string `json:"phoneNumber,omitempty"`

}

func (u *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref

	

	return json.Marshal(&struct { 
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		*Alias
	}{ 
		PhoneNumber: u.PhoneNumber,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
