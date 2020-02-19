package platformclientv2
import (
	"encoding/json"
)

// Contactlistfilterclause
type Contactlistfilterclause struct { 
	// FilterType - How to join predicates together.
	FilterType *string `json:"filterType,omitempty"`


	// Predicates - Conditions to filter the contacts by.
	Predicates *[]Contactlistfilterpredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactlistfilterclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
