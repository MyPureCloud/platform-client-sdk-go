package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdaymetriclisting
type Workdaymetriclisting struct { 
	// Entities
	Entities *[]Workdaymetric `json:"entities,omitempty"`

}

func (o *Workdaymetriclisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdaymetriclisting
	
	return json.Marshal(&struct { 
		Entities *[]Workdaymetric `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdaymetriclisting) UnmarshalJSON(b []byte) error {
	var WorkdaymetriclistingMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdaymetriclistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WorkdaymetriclistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdaymetriclisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
