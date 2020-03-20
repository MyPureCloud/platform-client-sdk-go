package platformclientv2
import (
	"encoding/json"
)

// Ntpsettings
type Ntpsettings struct { 
	// Servers - List of NTP servers, in priority order
	Servers *[]string `json:"servers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ntpsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
