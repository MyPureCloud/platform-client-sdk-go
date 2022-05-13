package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicsocialconversation
type Conversationsocialexpressioneventtopicsocialconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationsocialexpressioneventtopicsocialmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicsocialconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicsocialconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Conversationsocialexpressioneventtopicsocialmediaparticipant `json:"participants,omitempty"`
		
		OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		OtherMediaUris: o.OtherMediaUris,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicsocialconversation) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicsocialconversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicsocialconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationsocialexpressioneventtopicsocialconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationsocialexpressioneventtopicsocialconversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Participants, ok := ConversationsocialexpressioneventtopicsocialconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := ConversationsocialexpressioneventtopicsocialconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicsocialconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
