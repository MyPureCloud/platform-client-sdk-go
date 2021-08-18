package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangefromemailaddress
type Outboundmessagingmessagingcampaignconfigchangefromemailaddress struct { 
	// FriendlyName
	FriendlyName *string `json:"friendlyName,omitempty"`


	// LocalPart
	LocalPart *string `json:"localPart,omitempty"`


	// Domain
	Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`

}

func (u *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangefromemailaddress

	

	return json.Marshal(&struct { 
		FriendlyName *string `json:"friendlyName,omitempty"`
		
		LocalPart *string `json:"localPart,omitempty"`
		
		Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`
		*Alias
	}{ 
		FriendlyName: u.FriendlyName,
		
		LocalPart: u.LocalPart,
		
		Domain: u.Domain,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
