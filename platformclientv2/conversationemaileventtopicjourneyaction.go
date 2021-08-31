package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicjourneyaction
type Conversationemaileventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Conversationemaileventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Conversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationemaileventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var ConversationemaileventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationemaileventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationemaileventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := ConversationemaileventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
