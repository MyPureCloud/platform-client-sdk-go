package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulelistresponse
type Weekschedulelistresponse struct { 
	// Entities
	Entities *[]Weekschedulelistitemresponse `json:"entities,omitempty"`

}

func (o *Weekschedulelistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekschedulelistresponse
	
	return json.Marshal(&struct { 
		Entities *[]Weekschedulelistitemresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekschedulelistresponse) UnmarshalJSON(b []byte) error {
	var WeekschedulelistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekschedulelistresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WeekschedulelistresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedulelistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
