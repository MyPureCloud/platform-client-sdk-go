package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Outboundmessagingmessagingcampaignconfigchangeemailconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangeemailconfig

	

	return json.Marshal(&struct { 
		EmailColumns *[]string `json:"emailColumns,omitempty"`
		
		ContentTemplate *Outboundmessagingmessagingcampaignconfigchangeresponseref `json:"contentTemplate,omitempty"`
		
		FromAddress *Outboundmessagingmessagingcampaignconfigchangefromemailaddress `json:"fromAddress,omitempty"`
		
		ReplyToAddress *Outboundmessagingmessagingcampaignconfigchangereplytoemailaddress `json:"replyToAddress,omitempty"`
		*Alias
	}{ 
		EmailColumns: u.EmailColumns,
		
		ContentTemplate: u.ContentTemplate,
		
		FromAddress: u.FromAddress,
		
		ReplyToAddress: u.ReplyToAddress,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeemailconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
