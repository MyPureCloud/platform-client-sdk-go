package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings
type Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings struct { 
	// CallableWindows
	CallableWindows *[]Dialeroutboundsettingsconfigchangecallablewindow `json:"callableWindows,omitempty"`

}

func (o *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings
	
	return json.Marshal(&struct { 
		CallableWindows *[]Dialeroutboundsettingsconfigchangecallablewindow `json:"callableWindows,omitempty"`
		*Alias
	}{ 
		CallableWindows: o.CallableWindows,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangeautomatictimezonemappingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangeautomatictimezonemappingsettingsMap)
	if err != nil {
		return err
	}
	
	if CallableWindows, ok := DialeroutboundsettingsconfigchangeautomatictimezonemappingsettingsMap["callableWindows"].([]interface{}); ok {
		CallableWindowsString, _ := json.Marshal(CallableWindows)
		json.Unmarshal(CallableWindowsString, &o.CallableWindows)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
