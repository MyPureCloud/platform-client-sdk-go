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
	NumEvaluationsAssigned *int `json:"numEvaluationsAssigned,omitempty"`


	// NumEvaluationsStarted
	NumEvaluationsStarted *int `json:"numEvaluationsStarted,omitempty"`


	// NumEvaluationsCompleted
	NumEvaluationsCompleted *int `json:"numEvaluationsCompleted,omitempty"`


	// NumCalibrationsAssigned
	NumCalibrationsAssigned *int `json:"numCalibrationsAssigned,omitempty"`


	// NumCalibrationsStarted
	NumCalibrationsStarted *int `json:"numCalibrationsStarted,omitempty"`


	// NumCalibrationsCompleted
	NumCalibrationsCompleted *int `json:"numCalibrationsCompleted,omitempty"`


	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
