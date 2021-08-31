package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchessummaryresponse
type Shifttradematchessummaryresponse struct { 
	// Entities
	Entities *[]Weekshifttradematchessummaryresponse `json:"entities,omitempty"`

}

func (o *Shifttradematchessummaryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradematchessummaryresponse
	
	return json.Marshal(&struct { 
		Entities *[]Weekshifttradematchessummaryresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradematchessummaryresponse) UnmarshalJSON(b []byte) error {
	var ShifttradematchessummaryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradematchessummaryresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ShifttradematchessummaryresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradematchessummaryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
