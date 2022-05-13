package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sendagentlessoutboundmessagerequest
type Sendagentlessoutboundmessagerequest struct { 
	// FromAddress - The messaging address of the sender of the message. For an SMS messenger type, this must be a currently provisioned SMS phone number. For a WhatsApp messenger type use the provisioned WhatsApp integration’s ID
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234.
	ToAddress *string `json:"toAddress,omitempty"`


	// ToAddressMessengerType - The recipient messaging address messenger type.
	ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`


	// TextBody - The text of the message to send. This field is required in the case of SMS messenger type. Maximum character counts are: SMS - 765 characters, other channels - 2000 characters.
	TextBody *string `json:"textBody,omitempty"`


	// MessagingTemplate - The messaging template to use in the case of WhatsApp messenger type. This field is required when using WhatsApp messenger type
	MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`


	// UseExistingActiveConversation - Use an existing active conversation to send the agentless outbound message. Set this parameter to 'true' to use active conversation. Default value: false
	UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`

}

func (o *Sendagentlessoutboundmessagerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sendagentlessoutboundmessagerequest
	
	return json.Marshal(&struct { 
		FromAddress *string `json:"fromAddress,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`
		
		UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`
		*Alias
	}{ 
		FromAddress: o.FromAddress,
		
		ToAddress: o.ToAddress,
		
		ToAddressMessengerType: o.ToAddressMessengerType,
		
		TextBody: o.TextBody,
		
		MessagingTemplate: o.MessagingTemplate,
		
		UseExistingActiveConversation: o.UseExistingActiveConversation,
		Alias:    (*Alias)(o),
	})
}

func (o *Sendagentlessoutboundmessagerequest) UnmarshalJSON(b []byte) error {
	var SendagentlessoutboundmessagerequestMap map[string]interface{}
	err := json.Unmarshal(b, &SendagentlessoutboundmessagerequestMap)
	if err != nil {
		return err
	}
	
	if FromAddress, ok := SendagentlessoutboundmessagerequestMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if ToAddress, ok := SendagentlessoutboundmessagerequestMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if ToAddressMessengerType, ok := SendagentlessoutboundmessagerequestMap["toAddressMessengerType"].(string); ok {
		o.ToAddressMessengerType = &ToAddressMessengerType
	}
    
	if TextBody, ok := SendagentlessoutboundmessagerequestMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if MessagingTemplate, ok := SendagentlessoutboundmessagerequestMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	
	if UseExistingActiveConversation, ok := SendagentlessoutboundmessagerequestMap["useExistingActiveConversation"].(bool); ok {
		o.UseExistingActiveConversation = &UseExistingActiveConversation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
