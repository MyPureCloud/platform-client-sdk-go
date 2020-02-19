package platformclientv2
import (
	"encoding/json"
)

// Automatictimezonemappingsettings
type Automatictimezonemappingsettings struct { 
	// CallableWindows - The time intervals to use for automatic time zone mapping.
	CallableWindows *[]Callablewindow `json:"callableWindows,omitempty"`

}

// String returns a JSON representation of the model
func (o *Automatictimezonemappingsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
