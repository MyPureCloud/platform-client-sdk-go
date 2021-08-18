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


	// RouteId
	RouteId *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"routeId,omitempty"`

}

func (u *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress

	

	return json.Marshal(&struct { 
		Domain *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"domain,omitempty"`
		
		RouteId *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"routeId,omitempty"`
		*Alias
	}{ 
		Domain: u.Domain,
		
		RouteId: u.RouteId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
