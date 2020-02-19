package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicconversation
type Queueconversationeventtopicconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MaxParticipants
	MaxParticipants *int32 `json:"maxParticipants,omitempty"`


	// Participants
	Participants *[]Queueconversationeventtopicparticipant `json:"participants,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
