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


	// Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// User - Details of the user created the job
	User *Addressableentityref `json:"user,omitempty"`

}

func (u *Sendagentlessoutboundmessageresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sendagentlessoutboundmessageresponse

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		User *Addressableentityref `json:"user,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		ConversationId: u.ConversationId,
		
		FromAddress: u.FromAddress,
		
		ToAddress: u.ToAddress,
		
		MessengerType: u.MessengerType,
		
		TextBody: u.TextBody,
		
		MessagingTemplate: u.MessagingTemplate,
		
		Timestamp: Timestamp,
		
		SelfUri: u.SelfUri,
		
		User: u.User,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
