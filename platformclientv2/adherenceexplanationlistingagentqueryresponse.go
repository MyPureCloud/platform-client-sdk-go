package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationlistingagentqueryresponse
type Adherenceexplanationlistingagentqueryresponse struct { 
	// Entities
	Entities *[]Adherenceexplanationresponse `json:"entities,omitempty"`

}

func (o *Adherenceexplanationlistingagentqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationlistingagentqueryresponse
	
	return json.Marshal(&struct { 
		Entities *[]Adherenceexplanationresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationlistingagentqueryresponse) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationlistingagentqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationlistingagentqueryresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AdherenceexplanationlistingagentqueryresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationlistingagentqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
