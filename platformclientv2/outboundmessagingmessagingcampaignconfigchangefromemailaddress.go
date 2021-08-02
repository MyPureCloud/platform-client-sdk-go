package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangefromemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
