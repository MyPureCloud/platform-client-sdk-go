package platformclientv2
import (
	"encoding/json"
)

// Buschedulerunlisting
type Buschedulerunlisting struct { 
	// Entities
	Entities *[]Buschedulerun `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buschedulerunlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
