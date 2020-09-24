package platformclientv2
import (
	"encoding/json"
)

// Developmentactivityaggregatequeryrequestclause
type Developmentactivityaggregatequeryrequestclause struct { 
	// VarType - The logic used to combine the predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - The list of predicates used to filter the data
	Predicates *[]Developmentactivityaggregatequeryrequestpredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
