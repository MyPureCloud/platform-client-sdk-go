package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunitactivitycodelisting - List of BusinessUnitActivityCode
type Businessunitactivitycodelisting struct { 
	// Entities
	Entities *[]Businessunitactivitycode `json:"entities,omitempty"`

}

func (o *Businessunitactivitycodelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Businessunitactivitycodelisting
	
	return json.Marshal(&struct { 
		Entities *[]Businessunitactivitycode `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Businessunitactivitycodelisting) UnmarshalJSON(b []byte) error {
	var BusinessunitactivitycodelistingMap map[string]interface{}
	err := json.Unmarshal(b, &BusinessunitactivitycodelistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BusinessunitactivitycodelistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Businessunitactivitycodelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
