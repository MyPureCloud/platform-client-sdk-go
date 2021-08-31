package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopiccallconversation
type Queueconversationcalleventtopiccallconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Queueconversationcalleventtopiccallmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// MaxParticipants
	MaxParticipants *int `json:"maxParticipants,omitempty"`

}

func (o *Queueconversationcalleventtopiccallconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcalleventtopiccallconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Queueconversationcalleventtopiccallmediaparticipant `json:"participants,omitempty"`
		
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

func (o *Queueconversationcalleventtopiccallconversation) UnmarshalJSON(b []byte) error {
	var QueueconversationcalleventtopiccallconversationMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcalleventtopiccallconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcalleventtopiccallconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := QueueconversationcalleventtopiccallconversationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Participants, ok := QueueconversationcalleventtopiccallconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := QueueconversationcalleventtopiccallconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	
	if RecordingState, ok := QueueconversationcalleventtopiccallconversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
	
	if MaxParticipants, ok := QueueconversationcalleventtopiccallconversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopiccallconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
