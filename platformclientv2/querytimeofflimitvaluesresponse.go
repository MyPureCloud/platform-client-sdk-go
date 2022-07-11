package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Querytimeofflimitvaluesresponse
type Querytimeofflimitvaluesresponse struct { 
	// Values
	Values *[]Timeofflimitvaluerange `json:"values,omitempty"`

}

func (o *Querytimeofflimitvaluesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Querytimeofflimitvaluesresponse
	
	return json.Marshal(&struct { 
		Values *[]Timeofflimitvaluerange `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Querytimeofflimitvaluesresponse) UnmarshalJSON(b []byte) error {
	var QuerytimeofflimitvaluesresponseMap map[string]interface{}
	err := json.Unmarshal(b, &QuerytimeofflimitvaluesresponseMap)
	if err != nil {
		return err
	}
	
	if Values, ok := QuerytimeofflimitvaluesresponseMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Querytimeofflimitvaluesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
