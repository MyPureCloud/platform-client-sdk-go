package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicconversation
type Queueconversationsocialexpressioneventtopicconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MaxParticipants
	MaxParticipants *int32 `json:"maxParticipants,omitempty"`


	// Participants
	Participants *[]Queueconversationsocialexpressioneventtopicparticipant `json:"participants,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
