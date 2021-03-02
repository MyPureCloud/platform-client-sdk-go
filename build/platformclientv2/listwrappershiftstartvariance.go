package platformclientv2
import (
	"encoding/json"
)

// Listwrappershiftstartvariance
type Listwrappershiftstartvariance struct { 
	// Values
	Values *[]Shiftstartvariance `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listwrappershiftstartvariance) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
