package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sendagentlessoutboundmessageresponse
type Sendagentlessoutboundmessageresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ConversationId - The identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`


	// FromAddress - The sender of the message.
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The recipient of the message.
	ToAddress *string `json:"toAddress,omitempty"`


	// MessengerType - Type of messenger.
	MessengerType *string `json:"messengerType,omitempty"`


	// TextBody - The body of the text message.
	TextBody *string `json:"textBody,omitempty"`


	// MessagingTemplate - The messaging template sent
	MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`


	// UseExistingActiveConversation - Use an existing active conversation to send the agentless outbound message. Set this parameter to 'true' to use active conversation. Default value: false
	UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`


	// Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// User - Details of the user created the job
	User *Addressableentityref `json:"user,omitempty"`

}

func (o *Sendagentlessoutboundmessageresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sendagentlessoutboundmessageresponse
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`
		
		UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		User *Addressableentityref `json:"user,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ConversationId: o.ConversationId,
		
		FromAddress: o.FromAddress,
		
		ToAddress: o.ToAddress,
		
		MessengerType: o.MessengerType,
		
		TextBody: o.TextBody,
		
		MessagingTemplate: o.MessagingTemplate,
		
		UseExistingActiveConversation: o.UseExistingActiveConversation,
		
		Timestamp: Timestamp,
		
		SelfUri: o.SelfUri,
		
		User: o.User,
		Alias:    (*Alias)(o),
	})
}

func (o *Sendagentlessoutboundmessageresponse) UnmarshalJSON(b []byte) error {
	var SendagentlessoutboundmessageresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SendagentlessoutboundmessageresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SendagentlessoutboundmessageresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConversationId, ok := SendagentlessoutboundmessageresponseMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FromAddress, ok := SendagentlessoutboundmessageresponseMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if ToAddress, ok := SendagentlessoutboundmessageresponseMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if MessengerType, ok := SendagentlessoutboundmessageresponseMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
    
	if TextBody, ok := SendagentlessoutboundmessageresponseMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if MessagingTemplate, ok := SendagentlessoutboundmessageresponseMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	
	if UseExistingActiveConversation, ok := SendagentlessoutboundmessageresponseMap["useExistingActiveConversation"].(bool); ok {
		o.UseExistingActiveConversation = &UseExistingActiveConversation
	}
    
	if timestampString, ok := SendagentlessoutboundmessageresponseMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if SelfUri, ok := SendagentlessoutboundmessageresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if User, ok := SendagentlessoutboundmessageresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
