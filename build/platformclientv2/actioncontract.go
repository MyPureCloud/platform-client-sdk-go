package platformclientv2
import (
	"encoding/json"
)

// Actioncontract - This resource contains all of the schemas needed to define the inputs and outputs, of a single Action.
type Actioncontract struct { 
	// Output - The output to expect when executing this action.
	Output *Actionoutput `json:"output,omitempty"`


	// Input - The input required when executing this action.
	Input *Actioninput `json:"input,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actioncontract) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
