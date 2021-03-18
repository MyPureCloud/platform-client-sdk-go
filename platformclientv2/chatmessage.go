package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Chatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
