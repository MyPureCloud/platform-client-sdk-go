package platformclientv2
import (
	"encoding/json"
)

// Buschedulelisting
type Buschedulelisting struct { 
	// Entities
	Entities *[]Buschedulelistitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buschedulelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
