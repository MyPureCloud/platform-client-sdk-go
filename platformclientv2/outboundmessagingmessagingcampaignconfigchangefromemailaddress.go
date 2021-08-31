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

func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangefromemailaddress
	
	return json.Marshal(&struct { 
		FriendlyName *string `json:"friendlyName,omitempty"`
		
		LocalPart *string `json:"localPart,omitempty"`
		
		Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`
		*Alias
	}{ 
		FriendlyName: o.FriendlyName,
		
		LocalPart: o.LocalPart,
		
		Domain: o.Domain,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap)
	if err != nil {
		return err
	}
	
	if FriendlyName, ok := OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap["friendlyName"].(string); ok {
		o.FriendlyName = &FriendlyName
	}
	
	if LocalPart, ok := OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap["localPart"].(string); ok {
		o.LocalPart = &LocalPart
	}
	
	if Domain, ok := OutboundmessagingmessagingcampaignconfigchangefromemailaddressMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
