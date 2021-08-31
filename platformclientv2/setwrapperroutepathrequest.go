package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setwrapperroutepathrequest
type Setwrapperroutepathrequest struct { 
	// Values
	Values *[]Routepathrequest `json:"values,omitempty"`

}

func (o *Setwrapperroutepathrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setwrapperroutepathrequest
	
	return json.Marshal(&struct { 
		Values *[]Routepathrequest `json:"values,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Setwrapperroutepathrequest) UnmarshalJSON(b []byte) error {
	var SetwrapperroutepathrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SetwrapperroutepathrequestMap)
	if err != nil {
		return err
	}
	
	if Values, ok := SetwrapperroutepathrequestMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Setwrapperroutepathrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
