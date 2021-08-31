package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyaction
type Journeyaction struct { 
	// Id - The ID of an action from the Journey System (an action is spawned from an actionMap)
	Id *string `json:"id,omitempty"`


	// ActionMap - Details about the action map from the Journey System which triggered this action
	ActionMap *Journeyactionmap `json:"actionMap,omitempty"`

}

func (o *Journeyaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ActionMap *Journeyactionmap `json:"actionMap,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ActionMap: o.ActionMap,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyaction) UnmarshalJSON(b []byte) error {
	var JourneyactionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ActionMap, ok := JourneyactionMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
