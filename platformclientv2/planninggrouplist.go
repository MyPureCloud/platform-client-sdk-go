package platformclientv2
import (
	"encoding/json"
)

// Planninggrouplist - List of planning groups
type Planninggrouplist struct { 
	// Entities
	Entities *[]Planninggroup `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Planninggrouplist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
