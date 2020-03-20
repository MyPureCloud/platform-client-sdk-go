package platformclientv2
import (
	"encoding/json"
)

// Setwrapperdayofweek
type Setwrapperdayofweek struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Setwrapperdayofweek) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
