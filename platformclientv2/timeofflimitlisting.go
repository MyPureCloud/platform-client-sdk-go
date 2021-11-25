package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeofflimitlisting - The list of time off limit objects
type Timeofflimitlisting struct { 
	// Entities
	Entities *[]Timeofflimit `json:"entities,omitempty"`

}

func (o *Timeofflimitlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeofflimitlisting
	
	return json.Marshal(&struct { 
		Entities *[]Timeofflimit `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeofflimitlisting) UnmarshalJSON(b []byte) error {
	var TimeofflimitlistingMap map[string]interface{}
	err := json.Unmarshal(b, &TimeofflimitlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := TimeofflimitlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeofflimitlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
