package platformclientv2
import (
	"time"
	"encoding/json"
)

// Conversation
type Conversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// StartTime - The time when the conversation started. This will be the time when the first participant joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The time when the conversation ended. This will be the time when the last participant left the conversation, or null when the conversation is still active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndTime *time.Time `json:"endTime,omitempty"`


	// Address - The address of the conversation as seen from an external participant. For phone calls this will be the DNIS for inbound calls and the ANI for outbound calls. For other media types this will be the address of the destination participant for inbound and the address of the initiating participant for outbound.
	Address *string `json:"address,omitempty"`


	// Participants - The list of all participants in the conversation.
	Participants *[]Participant `json:"participants,omitempty"`


	// ConversationIds - A list of conversations to merge into this conversation to create a conference. This field is null except when being used to create a conference.
	ConversationIds *[]string `json:"conversationIds,omitempty"`


	// MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
	MaxParticipants *int32 `json:"maxParticipants,omitempty"`


	// RecordingState - On update, 'paused' initiates a secure pause, 'active' resumes any paused recordings; otherwise indicates state of conversation recording.
	RecordingState *string `json:"recordingState,omitempty"`


	// State - The conversation's state
	State *string `json:"state,omitempty"`


	// Divisions - Identifiers of divisions associated with this conversation
	Divisions *[]Conversationdivisionmembership `json:"divisions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
