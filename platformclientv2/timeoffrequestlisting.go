package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestlisting
type Timeoffrequestlisting struct { 
	// Entities - List of time off requests
	Entities *[]Timeoffrequest `json:"entities,omitempty"`

}

func (o *Timeoffrequestlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffrequestlisting
	
	return json.Marshal(&struct { 
		Entities *[]Timeoffrequest `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffrequestlisting) UnmarshalJSON(b []byte) error {
	var TimeoffrequestlistingMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := TimeoffrequestlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
