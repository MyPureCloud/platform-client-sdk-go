package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowdiagnosticinfo
type Flowdiagnosticinfo struct { 
	// LastActionId - The step number of the survey invite flow where the error occurred.
	LastActionId *int `json:"lastActionId,omitempty"`

}

func (o *Flowdiagnosticinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowdiagnosticinfo
	
	return json.Marshal(&struct { 
		LastActionId *int `json:"lastActionId,omitempty"`
		*Alias
	}{ 
		LastActionId: o.LastActionId,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowdiagnosticinfo) UnmarshalJSON(b []byte) error {
	var FlowdiagnosticinfoMap map[string]interface{}
	err := json.Unmarshal(b, &FlowdiagnosticinfoMap)
	if err != nil {
		return err
	}
	
	if LastActionId, ok := FlowdiagnosticinfoMap["lastActionId"].(float64); ok {
		LastActionIdInt := int(LastActionId)
		o.LastActionId = &LastActionIdInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowdiagnosticinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
