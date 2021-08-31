package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageconversation
type Messageconversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants - The list of participants involved in the conversation.
	Participants *[]Messagemediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris - The list of other media channels involved in the conversation.
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Messageconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messageconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Messagemediaparticipant `json:"participants,omitempty"`
		
		OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		OtherMediaUris: o.OtherMediaUris,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Messageconversation) UnmarshalJSON(b []byte) error {
	var MessageconversationMap map[string]interface{}
	err := json.Unmarshal(b, &MessageconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessageconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MessageconversationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Participants, ok := MessageconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := MessageconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	
	if SelfUri, ok := MessageconversationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messageconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
