package platformclientv2
import (
	"encoding/json"
)

// Dialogflowintent
type Dialogflowintent struct { 
	// Name - The intent name
	Name *string `json:"name,omitempty"`


	// Parameters - An object mapping parameter names to Parameter objects
	Parameters *map[string]Dialogflowparameter `json:"parameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialogflowintent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
