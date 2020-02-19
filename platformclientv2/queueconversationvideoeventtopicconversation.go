package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicconversation
type Queueconversationvideoeventtopicconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MaxParticipants
	MaxParticipants *int32 `json:"maxParticipants,omitempty"`


	// Participants
	Participants *[]Queueconversationvideoeventtopicparticipant `json:"participants,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
