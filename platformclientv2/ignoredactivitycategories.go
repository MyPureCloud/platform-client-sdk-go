package platformclientv2
import (
	"encoding/json"
)

// Ignoredactivitycategories
type Ignoredactivitycategories struct { 
	// Values - Activity categories list
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ignoredactivitycategories) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
