package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwardingeventcall
type Callforwardingeventcall struct { 
	// Targets
	Targets *[]Callforwardingeventtarget `json:"targets,omitempty"`

}

func (o *Callforwardingeventcall) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwardingeventcall
	
	return json.Marshal(&struct { 
		Targets *[]Callforwardingeventtarget `json:"targets,omitempty"`
		*Alias
	}{ 
		Targets: o.Targets,
		Alias:    (*Alias)(o),
	})
}

func (o *Callforwardingeventcall) UnmarshalJSON(b []byte) error {
	var CallforwardingeventcallMap map[string]interface{}
	err := json.Unmarshal(b, &CallforwardingeventcallMap)
	if err != nil {
		return err
	}
	
	if Targets, ok := CallforwardingeventcallMap["targets"].([]interface{}); ok {
		TargetsString, _ := json.Marshal(Targets)
		json.Unmarshal(TargetsString, &o.Targets)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callforwardingeventcall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
