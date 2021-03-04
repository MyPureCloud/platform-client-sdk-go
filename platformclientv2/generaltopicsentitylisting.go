package platformclientv2
import (
	"encoding/json"
)

// Generaltopicsentitylisting
type Generaltopicsentitylisting struct { 
	// Entities
	Entities *[]Generaltopic `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generaltopicsentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
