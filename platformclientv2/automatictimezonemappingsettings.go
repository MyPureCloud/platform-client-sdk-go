package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Automatictimezonemappingsettings
type Automatictimezonemappingsettings struct { 
	// CallableWindows - The time intervals to use for automatic time zone mapping.
	CallableWindows *[]Callablewindow `json:"callableWindows,omitempty"`

}

func (o *Automatictimezonemappingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Automatictimezonemappingsettings
	
	return json.Marshal(&struct { 
		CallableWindows *[]Callablewindow `json:"callableWindows,omitempty"`
		*Alias
	}{ 
		CallableWindows: o.CallableWindows,
		Alias:    (*Alias)(o),
	})
}

func (o *Automatictimezonemappingsettings) UnmarshalJSON(b []byte) error {
	var AutomatictimezonemappingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AutomatictimezonemappingsettingsMap)
	if err != nil {
		return err
	}
	
	if CallableWindows, ok := AutomatictimezonemappingsettingsMap["callableWindows"].([]interface{}); ok {
		CallableWindowsString, _ := json.Marshal(CallableWindows)
		json.Unmarshal(CallableWindowsString, &o.CallableWindows)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Automatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
