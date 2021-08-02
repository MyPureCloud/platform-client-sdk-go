package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangeemailconfig
type Outboundmessagingmessagingcampaignconfigchangeemailconfig struct { 
	// EmailColumns
	EmailColumns *[]string `json:"emailColumns,omitempty"`


	// ContentTemplate
	ContentTemplate *Outboundmessagingmessagingcampaignconfigchangeresponseref `json:"contentTemplate,omitempty"`


	// FromAddress
	FromAddress *Outboundmessagingmessagingcampaignconfigchangefromemailaddress `json:"fromAddress,omitempty"`


	// ReplyToAddress
	ReplyToAddress *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress `json:"replyToAddress,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeemailconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
