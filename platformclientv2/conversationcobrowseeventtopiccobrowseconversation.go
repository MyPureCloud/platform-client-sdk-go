package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopiccobrowseconversation
type Conversationcobrowseeventtopiccobrowseconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Conversationcobrowseeventtopiccobrowsemediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

func (o *Conversationcobrowseeventtopiccobrowseconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcobrowseeventtopiccobrowseconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Conversationcobrowseeventtopiccobrowsemediaparticipant `json:"participants,omitempty"`
		
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

func (o *Conversationcobrowseeventtopiccobrowseconversation) UnmarshalJSON(b []byte) error {
	var ConversationcobrowseeventtopiccobrowseconversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcobrowseeventtopiccobrowseconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcobrowseeventtopiccobrowseconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationcobrowseeventtopiccobrowseconversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Participants, ok := ConversationcobrowseeventtopiccobrowseconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if OtherMediaUris, ok := ConversationcobrowseeventtopiccobrowseconversationMap["otherMediaUris"].([]interface{}); ok {
		OtherMediaUrisString, _ := json.Marshal(OtherMediaUris)
		json.Unmarshal(OtherMediaUrisString, &o.OtherMediaUris)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopiccobrowseconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
