package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicjourneyaction
type Queueconversationvideoeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationvideoeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Queueconversationvideoeventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Queueconversationvideoeventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := QueueconversationvideoeventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
