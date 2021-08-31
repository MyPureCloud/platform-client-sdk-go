package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanlistresponse
type Workplanlistresponse struct { 
	// Entities
	Entities *[]Workplanlistitemresponse `json:"entities,omitempty"`

}

func (o *Workplanlistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanlistresponse
	
	return json.Marshal(&struct { 
		Entities *[]Workplanlistitemresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanlistresponse) UnmarshalJSON(b []byte) error {
	var WorkplanlistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanlistresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WorkplanlistresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanlistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
