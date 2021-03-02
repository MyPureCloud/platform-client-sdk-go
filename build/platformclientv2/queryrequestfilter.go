package platformclientv2
import (
	"encoding/json"
)

// Queryrequestfilter
type Queryrequestfilter struct { 
	// VarType - The logic used to combine the clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - The list of clauses used to filter the data
	Clauses *[]Queryrequestclause `json:"clauses,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryrequestfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
