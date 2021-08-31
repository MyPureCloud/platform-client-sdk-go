package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradepreviewresponse
type Shifttradepreviewresponse struct { 
	// Activities - List of activities that will make up the new shift if this shift trade is approved
	Activities *[]Shifttradeactivitypreviewresponse `json:"activities,omitempty"`

}

func (o *Shifttradepreviewresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradepreviewresponse
	
	return json.Marshal(&struct { 
		Activities *[]Shifttradeactivitypreviewresponse `json:"activities,omitempty"`
		*Alias
	}{ 
		Activities: o.Activities,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradepreviewresponse) UnmarshalJSON(b []byte) error {
	var ShifttradepreviewresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradepreviewresponseMap)
	if err != nil {
		return err
	}
	
	if Activities, ok := ShifttradepreviewresponseMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradepreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
