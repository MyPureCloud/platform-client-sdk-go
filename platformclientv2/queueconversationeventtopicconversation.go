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

func (o *Queueconversationeventtopicconversation) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		MaxParticipants: o.MaxParticipants,
		
		Participants: o.Participants,
		
		RecordingState: o.RecordingState,
		
		Address: o.Address,
		
		ExternalTag: o.ExternalTag,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicconversation) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicconversationMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationeventtopicconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MaxParticipants, ok := QueueconversationeventtopicconversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	
	if Participants, ok := QueueconversationeventtopicconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if RecordingState, ok := QueueconversationeventtopicconversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
	
	if Address, ok := QueueconversationeventtopicconversationMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if ExternalTag, ok := QueueconversationeventtopicconversationMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
