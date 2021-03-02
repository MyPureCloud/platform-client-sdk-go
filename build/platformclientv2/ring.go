package platformclientv2
import (
	"encoding/json"
)

// Ring
type Ring struct { 
	// ExpansionCriteria
	ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`


	// Actions
	Actions *Actions `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ring) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
