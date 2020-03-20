package platformclientv2
import (
	"encoding/json"
)

// Callablewindow
type Callablewindow struct { 
	// Mapped - The time interval to place outbound calls, for contacts that can be mapped to a time zone.
	Mapped *Atzmtimeslot `json:"mapped,omitempty"`


	// Unmapped - The time interval and time zone to place outbound calls, for contacts that cannot be mapped to a time zone.
	Unmapped *Atzmtimeslotwithtimezone `json:"unmapped,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callablewindow) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
