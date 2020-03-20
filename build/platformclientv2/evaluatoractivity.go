package platformclientv2
import (
	"encoding/json"
)

// Evaluatoractivity
type Evaluatoractivity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`


	// NumEvaluationsAssigned
	NumEvaluationsAssigned *int32 `json:"numEvaluationsAssigned,omitempty"`


	// NumEvaluationsStarted
	NumEvaluationsStarted *int32 `json:"numEvaluationsStarted,omitempty"`


	// NumEvaluationsCompleted
	NumEvaluationsCompleted *int32 `json:"numEvaluationsCompleted,omitempty"`


	// NumCalibrationsAssigned
	NumCalibrationsAssigned *int32 `json:"numCalibrationsAssigned,omitempty"`


	// NumCalibrationsStarted
	NumCalibrationsStarted *int32 `json:"numCalibrationsStarted,omitempty"`


	// NumCalibrationsCompleted
	NumCalibrationsCompleted *int32 `json:"numCalibrationsCompleted,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
