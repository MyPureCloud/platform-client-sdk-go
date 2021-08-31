package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicjourneyaction
type Conversationeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

func (o *Conversationeventtopicjourneyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicjourneyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Conversationeventtopicjourneyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicjourneyaction) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicjourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicjourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicjourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := ConversationeventtopicjourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
