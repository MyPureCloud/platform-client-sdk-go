package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationlistingbuqueryresponse
type Adherenceexplanationlistingbuqueryresponse struct { 
	// Entities
	Entities *[]Adherenceexplanationresponse `json:"entities,omitempty"`

}

func (o *Adherenceexplanationlistingbuqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationlistingbuqueryresponse
	
	return json.Marshal(&struct { 
		Entities *[]Adherenceexplanationresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationlistingbuqueryresponse) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationlistingbuqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationlistingbuqueryresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AdherenceexplanationlistingbuqueryresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationlistingbuqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
