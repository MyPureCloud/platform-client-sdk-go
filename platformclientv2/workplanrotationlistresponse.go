package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanrotationlistresponse
type Workplanrotationlistresponse struct { 
	// Entities
	Entities *[]Workplanrotationresponse `json:"entities,omitempty"`

}

func (o *Workplanrotationlistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanrotationlistresponse
	
	return json.Marshal(&struct { 
		Entities *[]Workplanrotationresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanrotationlistresponse) UnmarshalJSON(b []byte) error {
	var WorkplanrotationlistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanrotationlistresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WorkplanrotationlistresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanrotationlistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
