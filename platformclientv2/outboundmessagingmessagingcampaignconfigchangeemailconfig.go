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

func (o *Outboundmessagingmessagingcampaignconfigchangeemailconfig) MarshalJSON() ([]byte, error) {
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
		EmailColumns: o.EmailColumns,
		
		ContentTemplate: o.ContentTemplate,
		
		FromAddress: o.FromAddress,
		
		ReplyToAddress: o.ReplyToAddress,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangeemailconfig) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangeemailconfigMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangeemailconfigMap)
	if err != nil {
		return err
	}
	
	if EmailColumns, ok := OutboundmessagingmessagingcampaignconfigchangeemailconfigMap["emailColumns"].([]interface{}); ok {
		EmailColumnsString, _ := json.Marshal(EmailColumns)
		json.Unmarshal(EmailColumnsString, &o.EmailColumns)
	}
	
	if ContentTemplate, ok := OutboundmessagingmessagingcampaignconfigchangeemailconfigMap["contentTemplate"].(map[string]interface{}); ok {
		ContentTemplateString, _ := json.Marshal(ContentTemplate)
		json.Unmarshal(ContentTemplateString, &o.ContentTemplate)
	}
	
	if FromAddress, ok := OutboundmessagingmessagingcampaignconfigchangeemailconfigMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if ReplyToAddress, ok := OutboundmessagingmessagingcampaignconfigchangeemailconfigMap["replyToAddress"].(map[string]interface{}); ok {
		ReplyToAddressString, _ := json.Marshal(ReplyToAddress)
		json.Unmarshal(ReplyToAddressString, &o.ReplyToAddress)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeemailconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
