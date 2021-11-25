package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffplanlisting
type Timeoffplanlisting struct { 
	// Entities
	Entities *[]Timeoffplan `json:"entities,omitempty"`

}

func (o *Timeoffplanlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffplanlisting
	
	return json.Marshal(&struct { 
		Entities *[]Timeoffplan `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffplanlisting) UnmarshalJSON(b []byte) error {
	var TimeoffplanlistingMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffplanlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := TimeoffplanlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffplanlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
