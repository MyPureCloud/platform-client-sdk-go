package platformclientv2
import (
	"time"
	"encoding/json"
)

// Sipsearchpublicrequest
type Sipsearchpublicrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CallId - unique identification of the placed call
	CallId *string `json:"callId,omitempty"`


	// ToUser - SIP user to who the call was placed
	ToUser *string `json:"toUser,omitempty"`


	// FromUser - SIP user who placed the call
	FromUser *string `json:"fromUser,omitempty"`


	// ConversationId - Unique identification of the conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId - Unique identification of the participant
	ParticipantId *string `json:"participantId,omitempty"`


	// DateStart - Start date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateStart *time.Time `json:"dateStart,omitempty"`


	// DateEnd - End date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateEnd *time.Time `json:"dateEnd,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sipsearchpublicrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
