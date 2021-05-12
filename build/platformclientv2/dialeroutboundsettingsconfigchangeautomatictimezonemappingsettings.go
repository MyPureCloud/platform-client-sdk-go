package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings
type Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings struct { 
	// CallableWindows
	CallableWindows *[]Dialeroutboundsettingsconfigchangecallablewindow `json:"callableWindows,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
