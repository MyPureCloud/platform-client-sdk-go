package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress
type Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress struct { 
	// Domain
	Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`


	// Route
	Route *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"route,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress
	
	return json.Marshal(&struct { 
		Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`
		
		Route *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"route,omitempty"`
		*Alias
	}{ 
		Domain: o.Domain,
		
		Route: o.Route,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangereplytoemailaddressMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangereplytoemailaddressMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := OutboundmessagingmessagingcampaignconfigchangereplytoemailaddressMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	
	if Route, ok := OutboundmessagingmessagingcampaignconfigchangereplytoemailaddressMap["route"].(map[string]interface{}); ok {
		RouteString, _ := json.Marshal(Route)
		json.Unmarshal(RouteString, &o.Route)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
