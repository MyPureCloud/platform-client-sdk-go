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

func (u *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings

	

	return json.Marshal(&struct { 
		CallableWindows *[]Dialeroutboundsettingsconfigchangecallablewindow `json:"callableWindows,omitempty"`
		*Alias
	}{ 
		CallableWindows: u.CallableWindows,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
