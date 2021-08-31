package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradelistresponse
type Shifttradelistresponse struct { 
	// Entities
	Entities *[]Shifttraderesponse `json:"entities,omitempty"`

}

func (o *Shifttradelistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradelistresponse
	
	return json.Marshal(&struct { 
		Entities *[]Shifttraderesponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradelistresponse) UnmarshalJSON(b []byte) error {
	var ShifttradelistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradelistresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ShifttradelistresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradelistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
