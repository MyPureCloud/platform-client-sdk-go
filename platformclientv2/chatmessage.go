package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmessage
type Chatmessage struct { 
	// Body - The message body
	Body *string `json:"body,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// To - The message recipient
	To *string `json:"to,omitempty"`


	// From - The message sender
	From *string `json:"from,omitempty"`


	// Utc
	Utc *string `json:"utc,omitempty"`


	// Chat - The interaction id (if available)
	Chat *string `json:"chat,omitempty"`


	// Message - The message id
	Message *string `json:"message,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// BodyType - Type of the message body (v2 chats only)
	BodyType *string `json:"bodyType,omitempty"`


	// SenderCommunicationId - Communication of sender (v2 chats only)
	SenderCommunicationId *string `json:"senderCommunicationId,omitempty"`


	// ParticipantPurpose - Participant purpose of sender (v2 chats only)
	ParticipantPurpose *string `json:"participantPurpose,omitempty"`


	// User - The user information for the sender (if available)
	User *Chatmessageuser `json:"user,omitempty"`

}

func (o *Chatmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatmessage
	
	return json.Marshal(&struct { 
		Body *string `json:"body,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		To *string `json:"to,omitempty"`
		
		From *string `json:"from,omitempty"`
		
		Utc *string `json:"utc,omitempty"`
		
		Chat *string `json:"chat,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		BodyType *string `json:"bodyType,omitempty"`
		
		SenderCommunicationId *string `json:"senderCommunicationId,omitempty"`
		
		ParticipantPurpose *string `json:"participantPurpose,omitempty"`
		
		User *Chatmessageuser `json:"user,omitempty"`
		*Alias
	}{ 
		Body: o.Body,
		
		Id: o.Id,
		
		To: o.To,
		
		From: o.From,
		
		Utc: o.Utc,
		
		Chat: o.Chat,
		
		Message: o.Message,
		
		VarType: o.VarType,
		
		BodyType: o.BodyType,
		
		SenderCommunicationId: o.SenderCommunicationId,
		
		ParticipantPurpose: o.ParticipantPurpose,
		
		User: o.User,
		Alias:    (*Alias)(o),
	})
}

func (o *Chatmessage) UnmarshalJSON(b []byte) error {
	var ChatmessageMap map[string]interface{}
	err := json.Unmarshal(b, &ChatmessageMap)
	if err != nil {
		return err
	}
	
	if Body, ok := ChatmessageMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if Id, ok := ChatmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if To, ok := ChatmessageMap["to"].(string); ok {
		o.To = &To
	}
    
	if From, ok := ChatmessageMap["from"].(string); ok {
		o.From = &From
	}
    
	if Utc, ok := ChatmessageMap["utc"].(string); ok {
		o.Utc = &Utc
	}
    
	if Chat, ok := ChatmessageMap["chat"].(string); ok {
		o.Chat = &Chat
	}
    
	if Message, ok := ChatmessageMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if VarType, ok := ChatmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if BodyType, ok := ChatmessageMap["bodyType"].(string); ok {
		o.BodyType = &BodyType
	}
    
	if SenderCommunicationId, ok := ChatmessageMap["senderCommunicationId"].(string); ok {
		o.SenderCommunicationId = &SenderCommunicationId
	}
    
	if ParticipantPurpose, ok := ChatmessageMap["participantPurpose"].(string); ok {
		o.ParticipantPurpose = &ParticipantPurpose
	}
    
	if User, ok := ChatmessageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
