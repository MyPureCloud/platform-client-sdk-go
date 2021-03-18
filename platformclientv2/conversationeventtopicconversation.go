package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicconversation
type Conversationeventtopicconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MaxParticipants
	MaxParticipants *int `json:"maxParticipants,omitempty"`


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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
