package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Conversationcalleventtopiccallconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcalleventtopiccallconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Conversationcalleventtopiccallmediaparticipant `json:"participants,omitempty"`
		
		OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		OtherMediaUris: o.OtherMediaUris,
		
		RecordingState: o.RecordingState,
		
		MaxParticipants: o.MaxParticipants,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcalleventtopiccallconversation) UnmarshalJSON(b []byte) error {
	var ConversationcalleventtopiccallconversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcalleventtopiccallconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcalleventtopiccallconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationcalleventtopiccallconversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Participants, ok := ConversationcalleventtopiccallconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := ConversationcalleventtopiccallconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	
	if RecordingState, ok := ConversationcalleventtopiccallconversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    
	if MaxParticipants, ok := ConversationcalleventtopiccallconversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopiccallconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
