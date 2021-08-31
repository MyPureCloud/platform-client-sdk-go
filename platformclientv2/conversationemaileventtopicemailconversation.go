package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicemailconversation
type Conversationemaileventtopicemailconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationemaileventtopicemailmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

func (o *Conversationemaileventtopicemailconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicemailconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Conversationemaileventtopicemailmediaparticipant `json:"participants,omitempty"`
		
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

func (o *Conversationemaileventtopicemailconversation) UnmarshalJSON(b []byte) error {
	var ConversationemaileventtopicemailconversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationemaileventtopicemailconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationemaileventtopicemailconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ConversationemaileventtopicemailconversationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Participants, ok := ConversationemaileventtopicemailconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := ConversationemaileventtopicemailconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicemailconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
