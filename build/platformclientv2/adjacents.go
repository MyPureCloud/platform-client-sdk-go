package platformclientv2
import (
	"encoding/json"
)

// Adjacents
type Adjacents struct { 
	// Superiors
	Superiors *[]User `json:"superiors,omitempty"`


	// Siblings
	Siblings *[]User `json:"siblings,omitempty"`


	// DirectReports
	DirectReports *[]User `json:"directReports,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adjacents) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
