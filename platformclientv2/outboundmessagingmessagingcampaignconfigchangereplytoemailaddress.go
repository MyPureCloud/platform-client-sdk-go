package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
