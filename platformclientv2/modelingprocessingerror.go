package platformclientv2
import (
	"encoding/json"
)

// Modelingprocessingerror
type Modelingprocessingerror struct { 
	// InternalErrorCode - An internal code representing the type of error. ModelInputMissing for 'Model Builder inputs not found.' ModelInputInvalid for 'Model Builder inputs are invalid. Ensure the input data format is correct.' ModelFailed for 'An error occured while building the model with the given input.'
	InternalErrorCode *string `json:"internalErrorCode,omitempty"`


	// Description - A text description of the error
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Modelingprocessingerror) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
