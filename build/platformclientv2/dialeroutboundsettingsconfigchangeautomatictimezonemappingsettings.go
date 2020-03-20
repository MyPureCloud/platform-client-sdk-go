package platformclientv2
import (
	"encoding/json"
)

// Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings
type Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings struct { 
	// CallableWindows
	CallableWindows *[]Dialeroutboundsettingsconfigchangecallablewindow `json:"callableWindows,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
