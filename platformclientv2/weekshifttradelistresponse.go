package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekshifttradelistresponse
type Weekshifttradelistresponse struct { 
	// Entities
	Entities *[]Weekshifttraderesponse `json:"entities,omitempty"`

}

func (o *Weekshifttradelistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekshifttradelistresponse
	
	return json.Marshal(&struct { 
		Entities *[]Weekshifttraderesponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekshifttradelistresponse) UnmarshalJSON(b []byte) error {
	var WeekshifttradelistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekshifttradelistresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WeekshifttradelistresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekshifttradelistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
