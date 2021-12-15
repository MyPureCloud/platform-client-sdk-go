package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicjourneyaction - A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
type Queueconversationemaileventtopicjourneyaction struct { 
	// Id - The ID of an action from the Journey System (an action is spawned from an actionMap)
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Queueconversationemaileventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationemaileventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Queueconversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationemaileventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var QueueconversationemaileventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationemaileventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationemaileventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := QueueconversationemaileventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
