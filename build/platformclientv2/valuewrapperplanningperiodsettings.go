package platformclientv2
import (
	"encoding/json"
)

// Valuewrapperplanningperiodsettings - An object to provide context to nullable fields in PATCH requests
type Valuewrapperplanningperiodsettings struct { 
	// Value - The value for the associated field
	Value *Planningperiodsettings `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperplanningperiodsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
