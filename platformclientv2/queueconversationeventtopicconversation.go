package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicconversation
type Queueconversationeventtopicconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MaxParticipants
	MaxParticipants *int `json:"maxParticipants,omitempty"`


	// Participants
	Participants *[]Queueconversationeventtopicparticipant `json:"participants,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// ExternalTag
	ExternalTag *string `json:"externalTag,omitempty"`

}

func (u *Queueconversationeventtopicconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicconversation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		
		Participants *[]Queueconversationeventtopicparticipant `json:"participants,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		MaxParticipants: u.MaxParticipants,
		
		Participants: u.Participants,
		
		RecordingState: u.RecordingState,
		
		Address: u.Address,
		
		ExternalTag: u.ExternalTag,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
