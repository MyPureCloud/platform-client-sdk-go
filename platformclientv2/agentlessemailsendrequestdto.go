package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentlessemailsendrequestdto
type Agentlessemailsendrequestdto struct { 
	// SenderType - The direction of the message.
	SenderType *string `json:"senderType,omitempty"`


	// ConversationId - The identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`


	// FromAddress - The sender of the message.
	FromAddress *Emailaddress `json:"fromAddress,omitempty"`


	// ToAddresses - The recipient(s) of the message.
	ToAddresses *[]Emailaddress `json:"toAddresses,omitempty"`


	// ReplyToAddress - The address to use for reply.
	ReplyToAddress *Emailaddress `json:"replyToAddress,omitempty"`


	// Subject - The subject of the message.
	Subject *string `json:"subject,omitempty"`


	// TextBody - The Content of the message, in plain text.
	TextBody *string `json:"textBody,omitempty"`


	// HtmlBody - The Content of the message, in HTML. Links, images and styles are allowed
	HtmlBody *string `json:"htmlBody,omitempty"`

}

func (o *Agentlessemailsendrequestdto) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentlessemailsendrequestdto
	
	return json.Marshal(&struct { 
		SenderType *string `json:"senderType,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FromAddress *Emailaddress `json:"fromAddress,omitempty"`
		
		ToAddresses *[]Emailaddress `json:"toAddresses,omitempty"`
		
		ReplyToAddress *Emailaddress `json:"replyToAddress,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		HtmlBody *string `json:"htmlBody,omitempty"`
		*Alias
	}{ 
		SenderType: o.SenderType,
		
		ConversationId: o.ConversationId,
		
		FromAddress: o.FromAddress,
		
		ToAddresses: o.ToAddresses,
		
		ReplyToAddress: o.ReplyToAddress,
		
		Subject: o.Subject,
		
		TextBody: o.TextBody,
		
		HtmlBody: o.HtmlBody,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentlessemailsendrequestdto) UnmarshalJSON(b []byte) error {
	var AgentlessemailsendrequestdtoMap map[string]interface{}
	err := json.Unmarshal(b, &AgentlessemailsendrequestdtoMap)
	if err != nil {
		return err
	}
	
	if SenderType, ok := AgentlessemailsendrequestdtoMap["senderType"].(string); ok {
		o.SenderType = &SenderType
	}
    
	if ConversationId, ok := AgentlessemailsendrequestdtoMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FromAddress, ok := AgentlessemailsendrequestdtoMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if ToAddresses, ok := AgentlessemailsendrequestdtoMap["toAddresses"].([]interface{}); ok {
		ToAddressesString, _ := json.Marshal(ToAddresses)
		json.Unmarshal(ToAddressesString, &o.ToAddresses)
	}
	
	if ReplyToAddress, ok := AgentlessemailsendrequestdtoMap["replyToAddress"].(map[string]interface{}); ok {
		ReplyToAddressString, _ := json.Marshal(ReplyToAddress)
		json.Unmarshal(ReplyToAddressString, &o.ReplyToAddress)
	}
	
	if Subject, ok := AgentlessemailsendrequestdtoMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if TextBody, ok := AgentlessemailsendrequestdtoMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if HtmlBody, ok := AgentlessemailsendrequestdtoMap["htmlBody"].(string); ok {
		o.HtmlBody = &HtmlBody
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentlessemailsendrequestdto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
