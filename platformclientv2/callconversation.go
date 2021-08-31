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

func (o *Callconversation) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		OtherMediaUris: o.OtherMediaUris,
		
		RecordingState: o.RecordingState,
		
		MaxParticipants: o.MaxParticipants,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Callconversation) UnmarshalJSON(b []byte) error {
	var CallconversationMap map[string]interface{}
	err := json.Unmarshal(b, &CallconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := CallconversationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Participants, ok := CallconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := CallconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	
	if RecordingState, ok := CallconversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
	
	if MaxParticipants, ok := CallconversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	
	if SelfUri, ok := CallconversationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
