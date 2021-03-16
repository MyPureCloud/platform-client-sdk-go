package platformclientv2
import (
	"encoding/json"
)

// Workdaymetriclisting
type Workdaymetriclisting struct { 
	// Entities
	Entities *[]Workdaymetric `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workdaymetriclisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
