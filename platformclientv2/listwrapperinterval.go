package platformclientv2
import (
	"encoding/json"
)

// Listwrapperinterval
type Listwrapperinterval struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listwrapperinterval) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
