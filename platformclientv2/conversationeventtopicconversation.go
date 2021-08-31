package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// ExternalTag
	ExternalTag *string `json:"externalTag,omitempty"`

}

func (o *Conversationeventtopicconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		
		Participants *[]Conversationeventtopicparticipant `json:"participants,omitempty"`
		
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

func (o *Conversationeventtopicconversation) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicconversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MaxParticipants, ok := ConversationeventtopicconversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	
	if Participants, ok := ConversationeventtopicconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if RecordingState, ok := ConversationeventtopicconversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
	
	if Address, ok := ConversationeventtopicconversationMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if ExternalTag, ok := ConversationeventtopicconversationMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
