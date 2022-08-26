package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredgroups
type Unansweredgroups struct { 
	// Entities
	Entities *[]Unansweredgroup `json:"entities,omitempty"`

}

func (o *Unansweredgroups) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unansweredgroups
	
	return json.Marshal(&struct { 
		Entities *[]Unansweredgroup `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Unansweredgroups) UnmarshalJSON(b []byte) error {
	var UnansweredgroupsMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredgroupsMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := UnansweredgroupsMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredgroups) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
