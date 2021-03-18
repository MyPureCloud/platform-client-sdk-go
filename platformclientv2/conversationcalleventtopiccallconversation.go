package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopiccallconversation
type Conversationcalleventtopiccallconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationcalleventtopiccallmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// MaxParticipants
	MaxParticipants *int `json:"maxParticipants,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopiccallconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
