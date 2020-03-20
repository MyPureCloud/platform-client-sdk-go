package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicconversation
type Conversationeventtopicconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MaxParticipants
	MaxParticipants *int32 `json:"maxParticipants,omitempty"`


	// Participants
	Participants *[]Conversationeventtopicparticipant `json:"participants,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
