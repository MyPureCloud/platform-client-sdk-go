package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationlisting
type Adherenceexplanationlisting struct { 
	// Entities
	Entities *[]Adherenceexplanationresponse `json:"entities,omitempty"`

}

func (o *Adherenceexplanationlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationlisting
	
	return json.Marshal(&struct { 
		Entities *[]Adherenceexplanationresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationlisting) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationlistingMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AdherenceexplanationlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
