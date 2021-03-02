package platformclientv2
import (
	"encoding/json"
)

// Bushorttermforecastlisting
type Bushorttermforecastlisting struct { 
	// Entities
	Entities *[]Bushorttermforecastlistitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
