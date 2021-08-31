package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectioncontext
type Nludetectioncontext struct { 
	// Intent - Restrict detection to this intent.
	Intent *Contextintent `json:"intent,omitempty"`


	// Entity - Use this entity to restrict detection.
	Entity *Contextentity `json:"entity,omitempty"`

}

func (o *Nludetectioncontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectioncontext
	
	return json.Marshal(&struct { 
		Intent *Contextintent `json:"intent,omitempty"`
		
		Entity *Contextentity `json:"entity,omitempty"`
		*Alias
	}{ 
		Intent: o.Intent,
		
		Entity: o.Entity,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludetectioncontext) UnmarshalJSON(b []byte) error {
	var NludetectioncontextMap map[string]interface{}
	err := json.Unmarshal(b, &NludetectioncontextMap)
	if err != nil {
		return err
	}
	
	if Intent, ok := NludetectioncontextMap["intent"].(map[string]interface{}); ok {
		IntentString, _ := json.Marshal(Intent)
		json.Unmarshal(IntentString, &o.Intent)
	}
	
	if Entity, ok := NludetectioncontextMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nludetectioncontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
