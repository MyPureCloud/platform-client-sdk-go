package platformclientv2
import (
	"encoding/json"
)

// Developmentactivityaggregatequeryrequestpredicate
type Developmentactivityaggregatequeryrequestpredicate struct { 
	// Dimension - Each predicates specifies a dimension.
	Dimension *string `json:"dimension,omitempty"`


	// Value - Corresponding value for dimensions in predicates
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestpredicate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
