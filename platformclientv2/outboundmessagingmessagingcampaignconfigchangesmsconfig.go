package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangesmsconfig
type Outboundmessagingmessagingcampaignconfigchangesmsconfig struct { 
	// MessageColumn
	MessageColumn *string `json:"messageColumn,omitempty"`


	// PhoneColumn
	PhoneColumn *string `json:"phoneColumn,omitempty"`


	// SenderSmsPhoneNumber
	SenderSmsPhoneNumber *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref `json:"senderSmsPhoneNumber,omitempty"`


	// ContentTemplate
	ContentTemplate *Outboundmessagingmessagingcampaignconfigchangeresponseref `json:"contentTemplate,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangesmsconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangesmsconfig
	
	return json.Marshal(&struct { 
		MessageColumn *string `json:"messageColumn,omitempty"`
		
		PhoneColumn *string `json:"phoneColumn,omitempty"`
		
		SenderSmsPhoneNumber *Outboundmessagingmessagingcampaignconfigchangesmsphonenumberref `json:"senderSmsPhoneNumber,omitempty"`
		
		ContentTemplate *Outboundmessagingmessagingcampaignconfigchangeresponseref `json:"contentTemplate,omitempty"`
		*Alias
	}{ 
		MessageColumn: o.MessageColumn,
		
		PhoneColumn: o.PhoneColumn,
		
		SenderSmsPhoneNumber: o.SenderSmsPhoneNumber,
		
		ContentTemplate: o.ContentTemplate,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangesmsconfig) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangesmsconfigMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangesmsconfigMap)
	if err != nil {
		return err
	}
	
	if MessageColumn, ok := OutboundmessagingmessagingcampaignconfigchangesmsconfigMap["messageColumn"].(string); ok {
		o.MessageColumn = &MessageColumn
	}
	
	if PhoneColumn, ok := OutboundmessagingmessagingcampaignconfigchangesmsconfigMap["phoneColumn"].(string); ok {
		o.PhoneColumn = &PhoneColumn
	}
	
	if SenderSmsPhoneNumber, ok := OutboundmessagingmessagingcampaignconfigchangesmsconfigMap["senderSmsPhoneNumber"].(map[string]interface{}); ok {
		SenderSmsPhoneNumberString, _ := json.Marshal(SenderSmsPhoneNumber)
		json.Unmarshal(SenderSmsPhoneNumberString, &o.SenderSmsPhoneNumber)
	}
	
	if ContentTemplate, ok := OutboundmessagingmessagingcampaignconfigchangesmsconfigMap["contentTemplate"].(map[string]interface{}); ok {
		ContentTemplateString, _ := json.Marshal(ContentTemplate)
		json.Unmarshal(ContentTemplateString, &o.ContentTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangesmsconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
