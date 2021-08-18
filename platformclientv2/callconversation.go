package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callconversation
type Callconversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants - The list of participants involved in the conversation.
	Participants *[]Callmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris - The list of other media channels involved in the conversation.
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
	MaxParticipants *int `json:"maxParticipants,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Callconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callconversation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Callmediaparticipant `json:"participants,omitempty"`
		
		OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Participants: u.Participants,
		
		OtherMediaUris: u.OtherMediaUris,
		
		RecordingState: u.RecordingState,
		
		MaxParticipants: u.MaxParticipants,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
