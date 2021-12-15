package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicjourneyaction - A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
type Conversationsocialexpressioneventtopicjourneyaction struct { 
	// Id - The ID of an action from the Journey System (an action is spawned from an actionMap)
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationsocialexpressioneventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Conversationsocialexpressioneventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationsocialexpressioneventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := ConversationsocialexpressioneventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
