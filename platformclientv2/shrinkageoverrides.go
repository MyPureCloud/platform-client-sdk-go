package platformclientv2
import (
	"encoding/json"
)

// Shrinkageoverrides
type Shrinkageoverrides struct { 
	// Clear - Set true to clear the shrinkage interval overrides
	Clear *bool `json:"clear,omitempty"`


	// Values - List of interval shrinkage overrides
	Values *[]Shrinkageoverride `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shrinkageoverrides) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
