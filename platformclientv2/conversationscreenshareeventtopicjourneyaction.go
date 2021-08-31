package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicjourneyaction
type Conversationscreenshareeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationscreenshareeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Conversationscreenshareeventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationscreenshareeventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Conversationscreenshareeventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationscreenshareeventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var ConversationscreenshareeventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationscreenshareeventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationscreenshareeventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := ConversationscreenshareeventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
