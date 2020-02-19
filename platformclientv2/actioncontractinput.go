package platformclientv2
import (
	"encoding/json"
)

// Actioncontractinput - Contract definition.
type Actioncontractinput struct { 
	// Input - Execution input contract
	Input *Postinputcontract `json:"input,omitempty"`


	// Output - Execution output contract
	Output *Postoutputcontract `json:"output,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actioncontractinput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
