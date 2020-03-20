package platformclientv2
import (
	"encoding/json"
)

// Draftvalidationresult - Validation results
type Draftvalidationresult struct { 
	// Valid - Indicates if configuration is valid
	Valid *bool `json:"valid,omitempty"`


	// Errors - List of errors causing validation failure
	Errors *[]Errorbody `json:"errors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Draftvalidationresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
