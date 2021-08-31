package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Generaltopicsentitylisting
type Generaltopicsentitylisting struct { 
	// Entities
	Entities *[]Generaltopic `json:"entities,omitempty"`

}

func (o *Generaltopicsentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Generaltopicsentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Generaltopic `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Generaltopicsentitylisting) UnmarshalJSON(b []byte) error {
	var GeneraltopicsentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &GeneraltopicsentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := GeneraltopicsentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Generaltopicsentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
