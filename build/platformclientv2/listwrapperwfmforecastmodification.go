package platformclientv2
import (
	"encoding/json"
)

// Listwrapperwfmforecastmodification
type Listwrapperwfmforecastmodification struct { 
	// Values
	Values *[]Wfmforecastmodification `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listwrapperwfmforecastmodification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
