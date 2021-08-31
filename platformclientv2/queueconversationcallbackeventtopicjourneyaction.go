package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicjourneyaction
type Queueconversationcallbackeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationcallbackeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Queueconversationcallbackeventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Queueconversationcallbackeventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcallbackeventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := QueueconversationcallbackeventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
