package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Securesessionentitylisting
type Securesessionentitylisting struct { 
	// Entities
	Entities *[]Securesession `json:"entities,omitempty"`

}

func (o *Securesessionentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Securesessionentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Securesession `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Securesessionentitylisting) UnmarshalJSON(b []byte) error {
	var SecuresessionentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &SecuresessionentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := SecuresessionentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Securesessionentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
