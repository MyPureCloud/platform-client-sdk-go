package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Servicegoaltemplatelist - List of service goal templates
type Servicegoaltemplatelist struct { 
	// Entities
	Entities *[]Servicegoaltemplate `json:"entities,omitempty"`

}

func (o *Servicegoaltemplatelist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Servicegoaltemplatelist
	
	return json.Marshal(&struct { 
		Entities *[]Servicegoaltemplate `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Servicegoaltemplatelist) UnmarshalJSON(b []byte) error {
	var ServicegoaltemplatelistMap map[string]interface{}
	err := json.Unmarshal(b, &ServicegoaltemplatelistMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ServicegoaltemplatelistMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
