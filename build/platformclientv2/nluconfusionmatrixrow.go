package platformclientv2
import (
	"encoding/json"
)

// Nluconfusionmatrixrow
type Nluconfusionmatrixrow struct { 
	// Name - The name of the intent for the row.
	Name *string `json:"name,omitempty"`


	// Columns - The columns of confusion matrix for the intent
	Columns *[]Nluconfusionmatrixcolumn `json:"columns,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixrow) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
