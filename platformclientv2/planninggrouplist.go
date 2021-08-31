package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Planninggrouplist - List of planning groups
type Planninggrouplist struct { 
	// Entities
	Entities *[]Planninggroup `json:"entities,omitempty"`

}

func (o *Planninggrouplist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Planninggrouplist
	
	return json.Marshal(&struct { 
		Entities *[]Planninggroup `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Planninggrouplist) UnmarshalJSON(b []byte) error {
	var PlanninggrouplistMap map[string]interface{}
	err := json.Unmarshal(b, &PlanninggrouplistMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := PlanninggrouplistMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Planninggrouplist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
