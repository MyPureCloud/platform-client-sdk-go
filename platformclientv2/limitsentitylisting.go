package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Limitsentitylisting
type Limitsentitylisting struct { 
	// Entities
	Entities *[]Limit `json:"entities,omitempty"`

}

func (o *Limitsentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Limitsentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Limit `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Limitsentitylisting) UnmarshalJSON(b []byte) error {
	var LimitsentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &LimitsentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := LimitsentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Limitsentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
